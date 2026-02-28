// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"github.com/google/uuid"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

// CollectionGateway はコレクションのトランザクション処理インターフェース
type CollectionGateway interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Collection, error)
	List(ctx context.Context) ([]domain.Collection, error)
	Create(ctx context.Context, collection *domain.Collection) error
	Update(ctx context.Context, collection *domain.Collection) error
	Delete(ctx context.Context, collection *domain.Collection) error
}
