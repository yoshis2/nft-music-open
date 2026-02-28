// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"
	"log"
	"testing"

	"nft-music/domain"
	"nft-music/infrastructure/mysql"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupWalletTestDB() *WalletGateway {
	newMysql := mysql.NewTMysql()
	db := newMysql.TestOpen()

	err := db.AutoMigrate(&domain.Wallet{})
	if err != nil {
		log.Fatalln(err)
	}
	// 清掃
	db.Exec("DELETE FROM wallets")
	return &WalletGateway{Database: db}
}

func TestWallet_Get(t *testing.T) {
	walletRepo := setupWalletTestDB()
	address := "0xe178365cdd1ea2b6550217a1fde66dde54dad9cc3e11a4dbd6a806b739f05765"
	wallet := &domain.Wallet{
		ID:      uuid.New(),
		Address: address,
	}

	t.Run("正常系: アドレスで検索できる", func(t *testing.T) {
		err := walletRepo.Create(context.Background(), wallet)
		assert.NoError(t, err)

		result, err := walletRepo.Get(context.Background(), &domain.Wallet{Address: address})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, wallet.ID, result.ID)
		assert.Equal(t, address, result.Address)
	})

	t.Run("正常系: 存在しないアドレスの場合は空の構造体が返る", func(t *testing.T) {
		result, err := walletRepo.Get(context.Background(), &domain.Wallet{Address: "non-existent"})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, uuid.Nil, result.ID)
		assert.Equal(t, "", result.Address)
	})
}

func TestWallet_GetByID(t *testing.T) {
	walletRepo := setupWalletTestDB()
	id := uuid.New()
	address := "0x1234567890abcdef"
	wallet := &domain.Wallet{
		ID:      id,
		Address: address,
	}

	t.Run("正常系: IDで検索できる", func(t *testing.T) {
		err := walletRepo.Create(context.Background(), wallet)
		assert.NoError(t, err)

		result, err := walletRepo.GetByID(context.Background(), &domain.Wallet{ID: id})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, address, result.Address)
	})

	t.Run("正常系: 存在しないIDの場合は空の構造体が返る", func(t *testing.T) {
		result, err := walletRepo.GetByID(context.Background(), &domain.Wallet{ID: uuid.New()})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, uuid.Nil, result.ID)
	})
}

func TestWallet_List(t *testing.T) {
	walletRepo := setupWalletTestDB()

	id1 := uuid.New()
	id2 := uuid.New()

	wallet1 := &domain.Wallet{
		ID:      id1,
		Address: "0xa30442bd228d2081330238799a1c86aa4546977504eca63aaf9ab3099d7fd11c",
	}

	wallet2 := &domain.Wallet{
		ID:      id2,
		Address: "0xd6e8002014edebad8d5a7488e39bda6e7d9c808c79f96b8a03b40565d7885f1f",
	}

	err := walletRepo.Create(context.Background(), wallet1)
	assert.NoError(t, err)
	err = walletRepo.Create(context.Background(), wallet2)
	assert.NoError(t, err)

	wallets, err := walletRepo.List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, wallets, 2)
}

func TestWallet_Create(t *testing.T) {
	walletRepo := setupWalletTestDB()
	address := "0x54fca7608ff53d4103e37d366c0d8ef6592b1311865c631df56f344f85d8a272"
	wallet := &domain.Wallet{
		ID:      uuid.New(),
		Address: address,
	}

	err := walletRepo.Create(context.Background(), wallet)
	assert.NoError(t, err)

	result, err := walletRepo.Get(context.Background(), wallet)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, address, result.Address)
}
