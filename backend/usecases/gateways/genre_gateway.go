// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"github.com/google/uuid"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

// GenreGateway はジャンルのトラザクション処理インターフェース
type GenreGateway interface {
	Create(ctx context.Context, genre *domain.GenreMaster) error
	Get(ctx context.Context, id uuid.UUID) (*domain.GenreMaster, error)
	List(ctx context.Context) ([]domain.GenreMaster, error)
	Update(ctx context.Context, genre *domain.GenreMaster) error
	Delete(ctx context.Context, genre *domain.GenreMaster) error
}
