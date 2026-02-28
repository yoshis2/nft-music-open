// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"nft-music/contracts"
	"nft-music/domain"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmInteractor struct {
	Logging logging.Logging
}

func NewEvmInteractor(logging logging.Logging) *EvmInteractor {
	return &EvmInteractor{
		Logging: logging,
	}
}

func (interactor *EvmInteractor) Signer(ctx context.Context) (*ports.SignerOutput, error) {
	// 環境変数からGanacheのURLを取得。なければデフォルト値を使用
	ganacheURL := os.Getenv("GANACHE_URL")
	if ganacheURL == "" {
		ganacheURL = "http://ganache:8545" // GanacheのデフォルトRPCサーバー
	}
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA("01a3c379a8b07fb8f6ae7735fd7de35c5156c1cb8c3ac76a9d5157e3eedd2c4a")
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice := big.NewInt(20000000000) // 20 Gwei

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	name := "test"
	symbol := "tst"
	initialListingPrice := big.NewInt(domain.OneEther)
	initialRoyaltyFeeBps := big.NewInt(1000)

	address, tx, instance, err := contracts.DeployContracts(auth, client, name, symbol, initialListingPrice, initialRoyaltyFeeBps)
	if err != nil {
		panic(err)
	}

	fmt.Printf("instance: %v\n", instance)
	signerOutput := ports.SignerOutput{
		AddressHex:      address.Hex(),
		TransactionHash: tx.Hash().Hex(),
	}

	contractDial, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, err
	}

	log.Printf("contractDial : %v", &contractDial.ContractsCaller)
	return &signerOutput, nil
}
