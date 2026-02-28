// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"

	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/google/uuid"
)

// WalletInteractor は、ウォレット操作のユースケースを表します。
type WalletInteractor struct {
	Gateway gateways.WalletGateway
}

// NewWalletInteractor は、WalletInteractorを初期化します。
func NewWalletInteractor(gateway gateways.WalletGateway) *WalletInteractor {
	return &WalletInteractor{
		Gateway: gateway,
	}
}

// Get はWalletの存在を確認する
func (interactor *WalletInteractor) Get(ctx context.Context, input *ports.WalletInput) (*ports.WalletOutput, error) {
	wallet := &domain.Wallet{
		Address: input.Address,
	}

	output, err := interactor.Gateway.Get(ctx, wallet)
	if err != nil {
		return nil, err
	}

	if output == nil {
		return &ports.WalletOutput{}, nil
	}

	existWallet := &ports.WalletOutput{
		ID:        output.ID,
		Address:   output.Address,
		CreatedAt: output.CreatedAt,
	}

	return existWallet, nil
}

// List はウォレットアドレス情報を一覧で表示する
func (interactor *WalletInteractor) List(ctx context.Context) ([]ports.WalletOutput, error) {
	outputs, err := interactor.Gateway.List(ctx)
	if err != nil {
		return nil, err
	}

	var wallets []ports.WalletOutput
	for _, output := range outputs {
		var wallet ports.WalletOutput
		wallet.ID = output.ID
		wallet.Address = output.Address
		wallet.CreatedAt = output.CreatedAt

		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

// Create はWalletに登録する操作
func (interactor *WalletInteractor) Create(ctx context.Context, input *ports.WalletInput) (*ports.WalletOutput, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := util.JapaneseNowTime()

	wallet := &domain.Wallet{
		ID:        uuidV7,
		Address:   input.Address,
		CreatedAt: now,
	}

	if err := interactor.Gateway.Create(ctx, wallet); err != nil {
		return nil, err
	}

	output := &ports.WalletOutput{
		ID:        uuidV7,
		Address:   input.Address,
		CreatedAt: now,
	}

	return output, nil
}
