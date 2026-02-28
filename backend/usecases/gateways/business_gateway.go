// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"github.com/google/uuid"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

// BusinessGateway は職種のトランザクション処理インターフェース
type BusinessGateway interface {
	Create(ctx context.Context, business *domain.BusinessMaster) error
	Get(ctx context.Context, id uuid.UUID) (*domain.BusinessMaster, error)
	List(ctx context.Context) ([]domain.BusinessMaster, error)
	Update(ctx context.Context, business *domain.BusinessMaster) error
	Delete(ctx context.Context, business *domain.BusinessMaster) error
}
