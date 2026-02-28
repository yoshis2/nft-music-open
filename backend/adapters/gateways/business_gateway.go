// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"
	"errors"

	"nft-music/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BusinessGateway ビジネスマスターのリポジトリ
type BusinessGateway struct {
	Database *gorm.DB
}

func NewBusinessGateway(db *gorm.DB) *BusinessGateway {
	return &BusinessGateway{Database: db}
}

// Create は職種マスターの一つを追加する
func (gateway *BusinessGateway) Create(ctx context.Context, business *domain.BusinessMaster) error {
	return gateway.Database.WithContext(ctx).Create(&business).Error
}

// Get は特定の職種マスターを取得
func (gateway *BusinessGateway) Get(ctx context.Context, id uuid.UUID) (*domain.BusinessMaster, error) {
	var result domain.BusinessMaster
	if err := gateway.Database.WithContext(ctx).First(&result, "id = ?", &id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

// List は職種マスターの一覧を取得
func (gateway *BusinessGateway) List(ctx context.Context) ([]domain.BusinessMaster, error) {
	var result []domain.BusinessMaster
	if err := gateway.Database.WithContext(ctx).Order("updated_at desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// Update は職種マスターの一つを更新する
func (gateway *BusinessGateway) Update(ctx context.Context, business *domain.BusinessMaster) error {
	return gateway.Database.WithContext(ctx).Updates(&business).Error
}

// Delete は職種マスターの一つを削除する
func (gateway *BusinessGateway) Delete(ctx context.Context, business *domain.BusinessMaster) error {
	return gateway.Database.WithContext(ctx).Delete(business).Error
}
