// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"

	"gorm.io/gorm"
)

// TransactionGateway トランザクションリポジトリ
type TransactionGateway struct {
	Database *gorm.DB
}

func NewTransactionGateway(db *gorm.DB) *TransactionGateway {
	return &TransactionGateway{
		Database: db,
	}
}

func (gateway *TransactionGateway) List(ctx context.Context, limit int) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	// クエリのベースを作成
	query := gateway.Database.WithContext(ctx).Order("created_at DESC")

	// limitが0より大きい場合のみ、Limit句をクエリに追加
	if limit > 0 {
		query = query.Limit(limit)
	}

	// 最終的なクエリを実行
	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (gateway *TransactionGateway) ListByWallet(ctx context.Context, wallet string) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	if err := gateway.Database.WithContext(ctx).
		Joins("left join users on transactions.user_id = users.id").
		Where("`users`.`wallet` = ?", wallet).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (gateway *TransactionGateway) Search(ctx context.Context, genre string, minPrice int, maxPrice int, sort string) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction

	db := gateway.Database.WithContext(ctx)

	if genre != "" {
		db = db.Where("transactions.genre_id = ?", genre)
	}

	if minPrice > 0 {
		db = db.Where("transactions.price >= ?", minPrice)
	}

	if maxPrice > 0 {
		db = db.Where("transactions.price <= ?", maxPrice)
	}

	switch sort {
	case "price_asc":
		db = db.Order("price ASC")
	case "price_desc":
		db = db.Order("price DESC")
	default:
		db = db.Order("created_at DESC")
	}

	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (gateway *TransactionGateway) GetByTransactionid(ctx context.Context, transactionID string) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := gateway.Database.WithContext(ctx).
		Table("transactions").
		Select("transactions.*, genre_masters.name as genre_name").
		Joins("LEFT JOIN genre_masters ON transactions.genre_id = genre_masters.id").
		Where("transactions.id = ?", transactionID).
		First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (gateway *TransactionGateway) Create(ctx context.Context, transaction *domain.Transaction) error {
	return gateway.Database.WithContext(ctx).Create(&transaction).Error
}
