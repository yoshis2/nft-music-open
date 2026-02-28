// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"

	"nft-music/adapters/presenters"
	"nft-music/contracts"
	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator/v10"
)

type NftInteractor struct {
	UserGateway        gateways.UserGateway
	TransactionGateway gateways.TransactionGateway
	IpfsGateway        gateways.IpfsGateway
	EtherClient        *ethclient.Client
	Auth               *bind.TransactOpts
	Contracts          *contracts.Contracts
	Logging            logging.Logging
	Error              *presenters.ErrorPresenter
	Validator          *validator.Validate
}

func NewNftInteractor(userGateway gateways.UserGateway, transactionGateway gateways.TransactionGateway, ipfsGateway gateways.IpfsGateway, ethClient *ethclient.Client, auth *bind.TransactOpts, contracts *contracts.Contracts, logging logging.Logging, validate *validator.Validate) *NftInteractor {
	return &NftInteractor{
		UserGateway:        userGateway,
		TransactionGateway: transactionGateway,
		IpfsGateway:        ipfsGateway,
		EtherClient:        ethClient,
		Auth:               auth,
		Contracts:          contracts,
		Logging:            logging,
		Error:              presenters.NewErrorPresenter(logging),
		Validator:          validate,
	}
}

func (interactor *NftInteractor) List(ctx context.Context, limit int) ([]*ports.TransactionOutput, error) {
	outputs, err := interactor.TransactionGateway.List(ctx, limit)
	if err != nil {
		return nil, err
	}

	var transactions []*ports.TransactionOutput
	for _, output := range outputs {
		ipfsJSON, err := interactor.IpfsGateway.Get(ctx, output.TokenURL)
		if err != nil {
			return nil, err
		}

		transaction := outputPort(output, ipfsJSON)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (interactor *NftInteractor) ListByWallet(ctx context.Context, wallet string) ([]*ports.TransactionOutput, error) {
	outputs, err := interactor.TransactionGateway.ListByWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}

	var transactions []*ports.TransactionOutput
	for _, output := range outputs {
		ipfsJSON, err := interactor.IpfsGateway.Get(ctx, output.TokenURL)
		if err != nil {
			return nil, err
		}

		transaction := outputPort(output, ipfsJSON)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (interactor *NftInteractor) Search(ctx context.Context, query string, genre string, minPrice int, maxPrice int, sort string) ([]*ports.TransactionOutput, error) {
	outputs, err := interactor.TransactionGateway.Search(ctx, genre, minPrice, maxPrice, sort)
	if err != nil {
		return nil, err
	}

	var transactions []*ports.TransactionOutput
	for _, output := range outputs {
		ipfsJSON, err := interactor.IpfsGateway.Get(ctx, output.TokenURL)
		if err != nil {
			// 1つのNFTでエラーが発生しても処理を続行する
			interactor.Logging.Warning(fmt.Sprintf("failed to get ipfs json for token %s: %v", output.TokenURL, err))
			continue
		}

		// queryパラメータが存在する場合、IPFSメタJSONの内容でフィルタリングする
		if query != "" {
			searchQuery := util.NormalizeAndFold(query) // 検索クエリを正規化し、大文字小文字を区別しないようにする
			// ipfsJSON.NameとipfsJSON.Descriptionのいずれかに検索クエリが含まれているかチェック
			if !util.ContainsFold(util.NormalizeAndFold(ipfsJSON.Name), searchQuery) &&
				!util.ContainsFold(util.NormalizeAndFold(ipfsJSON.Description), searchQuery) {
				continue // 検索条件に合わない場合はスキップ
			}
		}

		transaction := outputPort(output, ipfsJSON)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (interactor *NftInteractor) GetByTransactionid(ctx context.Context, transactionID string) (*ports.TransactionOutput, error) {
	output, err := interactor.TransactionGateway.GetByTransactionid(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	ipfsJSON, err := interactor.IpfsGateway.Get(ctx, output.TokenURL)
	if err != nil {
		return nil, err
	}

	transaction := outputPort(output, ipfsJSON)

	return transaction, nil
}

func (interactor *NftInteractor) Mint(ctx context.Context, input *ports.NftInput, cid string) (*ports.TransactionOutput, error) {
	wallet := &domain.User{
		Wallet: input.Wallet,
	}
	user, err := interactor.UserGateway.GetByWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}

	price := big.NewInt(int64(input.Price))
	if input.ChainID == 1 || input.ChainID == 1337 || input.ChainID == 11155111 || input.ChainID == 5 || input.ChainID == 56 || input.ChainID == 97 || input.ChainID == 42161 || input.ChainID == 421613 || input.ChainID == 80001 {
		price = big.NewInt(int64(input.Price * 1000000000))
	}

	if price.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("price must be greater than 1")
	}

	header, err := interactor.EtherClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	// header.Number は Block Number
	nonce, err := interactor.EtherClient.NonceAt(ctx, interactor.Auth.From, header.Number)
	if err != nil {
		return nil, err
	}

	transactOpts := &bind.TransactOpts{
		From:     interactor.Auth.From,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   interactor.Auth.Signer,
		Context:  ctx,
		GasPrice: interactor.Auth.GasPrice,
		GasLimit: interactor.Auth.GasLimit,
		// Valueは呼び出しごとに設定する
	}

	// UpdateListingPriceはpayableではないためValueを0に設定
	transactOpts.Value = big.NewInt(0)
	updateTx, err := interactor.Contracts.UpdateListingPrice(transactOpts, price)
	if err != nil {
		interactor.Logging.Error(fmt.Sprintf("failed to update listing price: %s", err.Error()))
		return nil, err
	}
	// UpdateListingPrice トランザクションがブロックに書き込まれるのを待機します
	_, err = bind.WaitMined(ctx, interactor.EtherClient, updateTx)
	if err != nil {
		interactor.Logging.Error(fmt.Sprintf("failed to mine UpdateListingPrice transaction: %s", err.Error()))
		return nil, err
	}

	// これでコントラクトの `listingPrice` が `price` に更新されたことが保証されます。
	// そのため、`GetListingPrice` を呼び出す必要はなく、`CreateToken` に支払うValueは `price` となります。

	transactOpts.Nonce.Add(transactOpts.Nonce, big.NewInt(1))
	transactOpts.Value = price // CreateTokenではミント料として更新後のlistingPriceを設定
	trans, err := interactor.Contracts.CreateToken(transactOpts, fmt.Sprintf("https://ipfs.io/ipfs/%s", cid))
	if err != nil {
		interactor.Logging.Error(fmt.Sprintf("failed to create token: %s", err.Error()))
		return nil, err
	}

	floatPrice, _ := price.Float64()
	now := util.JapaneseNowTime()
	transactions := domain.Transaction{
		ID:        trans.Hash().Hex(),
		UserID:    user.ID,
		ChainID:   input.ChainID,
		Nonce:     int(trans.Nonce()),
		TokenURL:  fmt.Sprintf("/ipfs/%s", cid),
		GenreID:   input.GenreID,
		To:        sql.NullString{String: trans.To().Hex(), Valid: true},
		Price:     floatPrice,
		Insentive: input.Insentive,
		Cost:      int(trans.Cost().Int64()),
		Sale:      input.Sale,
		Status:    "created",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := interactor.TransactionGateway.Create(ctx, &transactions); err != nil {
		interactor.Logging.Error(fmt.Sprintf("failed to insert transaction: %s", err.Error()))
		return nil, err
	}

	return &ports.TransactionOutput{ // APIで返す構造体
		ID:          transactions.ID,
		UserID:      transactions.UserID,
		ChainID:     transactions.ChainID,
		Nonce:       transactions.Nonce,
		Name:        user.Name,
		Description: input.Description,
		ImageURL:    fmt.Sprintf("/ipfs/%s", cid),
		TokenURL:    transactions.TokenURL,
		GenreID:     transactions.GenreID,
		To:          transactions.To.String,
		Cost:        transactions.Cost,
		Status:      transactions.Status,
		Sale:        transactions.Sale,
		Price:       transactions.Price,
		Insentive:   transactions.Insentive,
		CreatedAt:   transactions.CreatedAt,
		UpdatedAt:   transactions.UpdatedAt,
	}, nil // APIで返す構造体
}

func outputPort(output *domain.Transaction, ipfsJSON *domain.IpfsJSON) *ports.TransactionOutput {
	price := output.Price
	if output.ChainID == 1 || output.ChainID == 1337 {
		price = output.Price / 1000000000
	}
	return &ports.TransactionOutput{
		ID:          output.ID,
		UserID:      output.UserID,
		ChainID:     output.ChainID,
		Nonce:       output.Nonce,
		Name:        ipfsJSON.Name,
		Description: ipfsJSON.Description,
		FileType:    ipfsJSON.FileType,
		ImageURL:    fmt.Sprintf("/ipfs/%s", ipfsJSON.ImageCid), // ipfsJSON.Cid,
		AudioURL:    fmt.Sprintf("/ipfs/%s", ipfsJSON.AudioCid),
		VideoURL:    fmt.Sprintf("/ipfs/%s", ipfsJSON.VideoCid),
		TokenURL:    output.TokenURL,
		GenreID:     output.GenreID,
		To:          output.To.String,
		Price:       price,
		Insentive:   output.Insentive,
		Cost:        output.Cost,
		Sale:        output.Sale,
		Status:      output.Status,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	}
}
