// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"nft-music/domain"
	"nft-music/infrastructure/mysql"
	"nft-music/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// TestMain は、パッケージ内の全テストの実行前に一度だけ呼び出される特別な関数です。
// データベース接続などの重いセットアップ処理をここで行います。
func TestMain(m *testing.M) {
	// データベースへの接続を一度だけ行う
	newMysql := mysql.NewTMysql()
	db := newMysql.TestOpen()
	testDB = db

	// 全てのテストを実行
	code := m.Run()

	// 後片付け処理があればここに記述
	// (今回は t.Cleanup を使うため、ここでは不要)

	os.Exit(code)
}

// setupBusinessTestDB は、各テストの実行前に呼び出され、
// データベースの状態をクリーンにして準備します。
func setupBusinessTestDB(t *testing.T) *BusinessGateway {
	// テスト終了後にテーブルをドロップしてクリーンアップするよう登録
	t.Cleanup(func() {
		err := testDB.Migrator().DropTable(&domain.BusinessMaster{})
		if err != nil {
			log.Fatalf("Failed to drop table: %v", err)
		}
	})

	// 各テストの開始前にテーブルをマイグレーション
	err := testDB.AutoMigrate(&domain.BusinessMaster{})
	// require を使うと、アサーションが失敗した場合にテストを即時中断できる
	require.NoError(t, err, "Failed to migrate database")

	return &BusinessGateway{Database: testDB}
}

func TestBusiness_Create(t *testing.T) {
	t.Run("success: 新しい職種を一件作成できる", func(t *testing.T) {
		// ARRANGE: テストの準備
		businessRepo := setupBusinessTestDB(t)
		now := util.JapaneseNowTime()
		business1 := &domain.BusinessMaster{
			ID:        uuid.New(),
			Name:      "新しい職種",
			CreatedAt: now,
			UpdatedAt: now,
		}

		// ACT: 実際の処理を実行
		err := businessRepo.Create(context.Background(), business1)

		// ASSERT: 結果を検証
		assert.NoError(t, err)
	})
}

func TestBusiness_Get(t *testing.T) {
	t.Run("success: 作成した職種を一件取得できる", func(t *testing.T) {
		// ARRANGE
		businessRepo := setupBusinessTestDB(t)
		now := util.JapaneseNowTime()
		business1 := &domain.BusinessMaster{
			ID:        uuid.New(),
			Name:      "取得テスト用の職種",
			CreatedAt: now,
			UpdatedAt: now,
		}
		// 最初にテスト用のデータを作成
		err := businessRepo.Create(context.Background(), business1)
		require.NoError(t, err, "Setup for Get test failed")

		// ACT
		result, err := businessRepo.Get(context.Background(), business1.ID)

		// ASSERT
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, business1.Name, result.Name)
	})
}

func TestBusiness_List(t *testing.T) {
	t.Run("success: 作成した全ての職種を一覧取得できる", func(t *testing.T) {
		// ARRANGE
		businessRepo := setupBusinessTestDB(t)
		now := time.Now()
		business1 := &domain.BusinessMaster{ID: uuid.New(), Name: "職種1", CreatedAt: now, UpdatedAt: now}
		business2 := &domain.BusinessMaster{ID: uuid.New(), Name: "職種2", CreatedAt: now, UpdatedAt: now}

		// 複数のデータを作成
		require.NoError(t, businessRepo.Create(context.Background(), business1))
		require.NoError(t, businessRepo.Create(context.Background(), business2))

		// ACT
		results, err := businessRepo.List(context.Background())

		// ASSERT
		assert.NoError(t, err)
		assert.Equal(t, 2, len(results))
	})
}

func TestBusiness_Update(t *testing.T) {
	t.Run("success: 職種の名前を変更できる", func(t *testing.T) {
		// ARRANGE
		businessRepo := setupBusinessTestDB(t)
		now := time.Now()
		businessToUpdate := &domain.BusinessMaster{
			ID:        uuid.New(),
			Name:      "変更前の職種",
			CreatedAt: now,
			UpdatedAt: now,
		}
		require.NoError(t, businessRepo.Create(context.Background(), businessToUpdate))

		updatedBusiness := &domain.BusinessMaster{
			ID:   businessToUpdate.ID,
			Name: "変更された職種",
		}

		// ACT
		err := businessRepo.Update(context.Background(), updatedBusiness)
		require.NoError(t, err)

		// ASSERT
		result, err := businessRepo.Get(context.Background(), businessToUpdate.ID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "変更された職種", result.Name)
		// 更新日時が変わっていることなども検証できるとより良い
	})
}

func TestBusiness_Delete(t *testing.T) {
	t.Run("success: 作成した職種を削除できる", func(t *testing.T) {
		// ARRANGE
		businessRepo := setupBusinessTestDB(t)
		now := time.Now()
		business1 := &domain.BusinessMaster{ID: uuid.New(), Name: "削除される職種1", CreatedAt: now, UpdatedAt: now}
		business2 := &domain.BusinessMaster{ID: uuid.New(), Name: "残る職種2", CreatedAt: now, UpdatedAt: now}
		require.NoError(t, businessRepo.Create(context.Background(), business1))
		require.NoError(t, businessRepo.Create(context.Background(), business2))

		// ACT
		err := businessRepo.Delete(context.Background(), &domain.BusinessMaster{ID: business1.ID})

		// ASSERT
		assert.NoError(t, err)
		// 削除されていることを確認
		result, err := businessRepo.Get(context.Background(), business1.ID)
		assert.NoError(t, err)
		assert.Nil(t, result)

		// 削除されていないデータが残っていることを確認
		results, err := businessRepo.List(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, len(results))
		assert.Equal(t, business2.Name, results[0].Name)
	})
}
