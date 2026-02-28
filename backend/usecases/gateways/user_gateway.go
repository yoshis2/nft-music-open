// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

type UserGateway interface {
	Get(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByWallet(ctx context.Context, user *domain.User) (*domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, user *domain.User) error
}
