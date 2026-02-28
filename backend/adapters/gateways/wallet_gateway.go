// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"
	"errors"

	"nft-music/domain"

	"gorm.io/gorm"
)

// WalletGateway ウォレット リポジトリ
type WalletGateway struct {
	Database *gorm.DB
}

func NewWalletGateway(db *gorm.DB) *WalletGateway {
	return &WalletGateway{Database: db}
}

// Get 指定したウォレットアドレスを抽出する
func (gateway *WalletGateway) Get(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	var result domain.Wallet

	if err := gateway.Database.WithContext(ctx).First(&result, "address = ?", wallet.Address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.Wallet{}, nil
		}
		return nil, err
	}
	return &result, nil
}

func (gateway *WalletGateway) GetByID(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	var result domain.Wallet

	if err := gateway.Database.WithContext(ctx).First(&result, "id = ?", wallet.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.Wallet{}, nil
		}
		return nil, err
	}
	return &result, nil
}

// List 指定したウォレットアドレスを抽出する
func (gateway *WalletGateway) List(ctx context.Context) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	if err := gateway.Database.WithContext(ctx).Order("created_at desc").Limit(200).Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}

// Create ウォレットDBにデータを挿入
func (gateway *WalletGateway) Create(ctx context.Context, wallet *domain.Wallet) error {
	return gateway.Database.WithContext(ctx).Create(&wallet).Error
}
