// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

// WalletGateway インターフェースはDB処理になる
type WalletGateway interface {
	Get(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error)
	GetByID(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error)
	List(ctx context.Context) ([]domain.Wallet, error)
	Create(ctx context.Context, wallet *domain.Wallet) error
}
