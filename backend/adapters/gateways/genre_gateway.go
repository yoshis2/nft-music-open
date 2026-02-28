// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenreGateway ジャンルマスターの構造体
type GenreGateway struct {
	Database *gorm.DB
}

func NewGenreGateway(db *gorm.DB) *GenreGateway {
	return &GenreGateway{Database: db}
}

// Create はジャンルを一つ追加する
func (gateway *GenreGateway) Create(ctx context.Context, genre *domain.GenreMaster) error {
	return gateway.Database.WithContext(ctx).Create(&genre).Error
}

// Get は指定した一つのジャンルを取得
func (gateway *GenreGateway) Get(ctx context.Context, id uuid.UUID) (*domain.GenreMaster, error) {
	var result domain.GenreMaster
	if err := gateway.Database.WithContext(ctx).First(&result, "id = ?", &id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// List はジャンルの一覧を取得
func (gateway *GenreGateway) List(ctx context.Context) ([]domain.GenreMaster, error) {
	var results []domain.GenreMaster
	if err := gateway.Database.WithContext(ctx).Order("updated_at desc").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// Update はジャンルを一つ編集する
func (gateway *GenreGateway) Update(ctx context.Context, genre *domain.GenreMaster) error {
	return gateway.Database.WithContext(ctx).Updates(&genre).Error
}

// Delete はジャンルを一つ削除する
func (gateway *GenreGateway) Delete(ctx context.Context, genre *domain.GenreMaster) error {
	return gateway.Database.WithContext(ctx).Delete(genre).Error
}
