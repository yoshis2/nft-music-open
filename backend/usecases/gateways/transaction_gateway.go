// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"

	"nft-music/domain"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE

// TransactionGateway はトランザクションのトランザクション処理インターフェース
type TransactionGateway interface {
	List(ctx context.Context, limit int) ([]*domain.Transaction, error)
	ListByWallet(ctx context.Context, wallet string) ([]*domain.Transaction, error)
	Search(ctx context.Context, genre string, minPrice int, maxPrice int, sort string) ([]*domain.Transaction, error)
	GetByTransactionid(ctx context.Context, transactionID string) (*domain.Transaction, error)
	Create(ctx context.Context, transaction *domain.Transaction) error
}
