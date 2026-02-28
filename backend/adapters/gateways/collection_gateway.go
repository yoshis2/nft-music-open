// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CollectionGateway struct {
	Database *gorm.DB
}

func NewCollectionGateway(db *gorm.DB) *CollectionGateway {
	return &CollectionGateway{
		Database: db,
	}
}

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

func (gateway *CollectionGateway) Get(ctx context.Context, id uuid.UUID) (*domain.Collection, error) {
	var result domain.Collection
	if err := gateway.Database.WithContext(ctx).First(&result, "id = ?", &id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (gateway *CollectionGateway) List(ctx context.Context) ([]domain.Collection, error) {
	var result []domain.Collection
	if err := gateway.Database.WithContext(ctx).Order("updated_at desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (gateway *CollectionGateway) Create(ctx context.Context, collection *domain.Collection) error {
	return gateway.Database.WithContext(ctx).Create(&collection).Error
}

func (gateway *CollectionGateway) Update(ctx context.Context, collection *domain.Collection) error {
	return gateway.Database.WithContext(ctx).Updates(&collection).Error
}

func (gateway *CollectionGateway) Delete(ctx context.Context, collection *domain.Collection) error {
	return gateway.Database.WithContext(ctx).Delete(collection).Error
}
