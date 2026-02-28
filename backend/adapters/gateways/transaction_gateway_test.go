// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"nft-music/domain"
	"nft-music/infrastructure/mysql"
	"nft-music/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupTransactionTestDB() *TransactionGateway {
	if db == nil {
		newMysql := mysql.NewTMysql()
		db = newMysql.TestOpen()
	}

	// Drop tables if they exist to ensure a clean slate
	if err := db.Migrator().DropTable(&domain.Transaction{}, &domain.User{}, &domain.GenreMaster{}); err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}

	// AutoMigrate will create tables, columns, and indexes
	if err := db.AutoMigrate(&domain.User{}, &domain.GenreMaster{}, &domain.Transaction{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Truncate tables to ensure clean state for each test run
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE transactions")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("TRUNCATE TABLE genre_masters")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	return &TransactionGateway{Database: db}
}

func seedData() {
	now := util.JapaneseNowTime()
	userID1, _ := uuid.NewRandom()
	userID2, _ := uuid.NewRandom()
	genreID1, _ := uuid.NewRandom()
	genreID2, _ := uuid.NewRandom()
	ipfsID1 := "ipfs_id_1"
	ipfsID2 := "ipfs_id_2"
	ipfsID3 := "ipfs_id_3"

	// --- Seed Users ---
	db.Create(&domain.User{ID: userID1, Wallet: "wallet1", Name: "user1", CreatedAt: now, UpdatedAt: now})
	db.Create(&domain.User{ID: userID2, Wallet: "wallet2", Name: "user2", CreatedAt: now, UpdatedAt: now})

	// --- Seed Genres ---
	db.Create(&domain.GenreMaster{ID: genreID1, Name: "Rock", CreatedAt: now, UpdatedAt: now})
	db.Create(&domain.GenreMaster{ID: genreID2, Name: "Pop", CreatedAt: now, UpdatedAt: now})

	// --- Seed Transactions ---
	db.Create(&domain.Transaction{ID: "tx1", UserID: userID1, TokenURL: fmt.Sprintf("/ipfs/%s", ipfsID1), GenreID: genreID1, Price: 100, CreatedAt: now.Add(-time.Hour * 2), UpdatedAt: now})
	db.Create(&domain.Transaction{ID: "tx2", UserID: userID2, TokenURL: fmt.Sprintf("/ipfs/%s", ipfsID2), GenreID: genreID2, Price: 200, CreatedAt: now.Add(-time.Hour * 1), UpdatedAt: now})
	db.Create(&domain.Transaction{ID: "tx3", UserID: userID1, TokenURL: fmt.Sprintf("/ipfs/%s", ipfsID3), GenreID: genreID1, Price: 150, CreatedAt: now, UpdatedAt: now})
}

func TestTransactionGateway_Search(t *testing.T) {
	gateway := setupTransactionTestDB()
	seedData()
	ctx := context.Background()

	t.Run("by genre ID", func(t *testing.T) {
		var genre PopGenreMaster
		db.Where("name = ?", "Pop").First(&genre)
		results, err := gateway.Search(ctx, genre.ID.String(), 0, 0, "")
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "tx2", results[0].ID)
	})

	t.Run("by min price 150", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 150, 0, "")
		assert.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("by max price 150", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 0, 150, "")
		assert.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("by price range 120 to 180", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 120, 180, "")
		assert.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "tx3", results[0].ID)
	})

	t.Run("sort by price asc", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 0, 0, "price_asc")
		assert.NoError(t, err)
		assert.Len(t, results, 3)
		assert.Equal(t, "tx1", results[0].ID)
		assert.Equal(t, "tx3", results[1].ID)
		assert.Equal(t, "tx2", results[2].ID)
	})

	t.Run("sort by price desc", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 0, 0, "price_desc")
		assert.NoError(t, err)
		assert.Len(t, results, 3)
		assert.Equal(t, "tx2", results[0].ID)
		assert.Equal(t, "tx3", results[1].ID)
		assert.Equal(t, "tx1", results[2].ID)
	})

	t.Run("sort by newest (default)", func(t *testing.T) {
		results, err := gateway.Search(ctx, "", 0, 0, "")
		assert.NoError(t, err)
		assert.Len(t, results, 3)
		assert.Equal(t, "tx3", results[0].ID)
		assert.Equal(t, "tx2", results[1].ID)
		assert.Equal(t, "tx1", results[2].ID)
	})
}

type PopGenreMaster struct {
	ID uuid.UUID `gorm:"primaryKey;type:char(36)"`
}

func (PopGenreMaster) TableName() string {
	return "genre_masters"
}

type RockGenreMaster struct {
	ID uuid.UUID `gorm:"primaryKey;type:char(36)"`
}

func (RockGenreMaster) TableName() string {
	return "genre_masters"
}
