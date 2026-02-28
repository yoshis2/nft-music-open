// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"gorm.io/gorm"
)

type UserGateway struct {
	Database *gorm.DB
}

func NewUserGateway(db *gorm.DB) *UserGateway {
	return &UserGateway{Database: db}
}

func (gateway *UserGateway) Get(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := gateway.Database.WithContext(ctx).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (gateway *UserGateway) GetByWallet(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := gateway.Database.WithContext(ctx).Where("wallet = ?", user.Wallet).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (gateway *UserGateway) List(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := gateway.Database.WithContext(ctx).Order("updated_at desc").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (gateway *UserGateway) Create(ctx context.Context, user *domain.User) error {
	return gateway.Database.WithContext(ctx).Create(&user).Error
}

func (gateway *UserGateway) Update(ctx context.Context, user *domain.User) error {
	return gateway.Database.WithContext(ctx).Updates(&user).Error
}

func (gateway *UserGateway) Delete(ctx context.Context, user *domain.User) error {
	return gateway.Database.WithContext(ctx).Delete(user).Error
}
