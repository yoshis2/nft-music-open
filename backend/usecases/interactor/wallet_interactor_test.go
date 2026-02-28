// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"nft-music/domain"
	"nft-music/usecases/gateways/mock"
	"nft-music/usecases/ports"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWalletInteractor_Get(t *testing.T) {
	id := uuid.New()
	address := "0x884A3d4a6a09be8EAC4e0e2E5CfbbE968f63B448"
	now := time.Now()

	t.Run("正常系: ウォレット情報を取得できる", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockGateway := mock.NewMockWalletGateway(ctrl)
		interactor := NewWalletInteractor(mockGateway)

		expectedWallet := &domain.Wallet{
			ID:        id,
			Address:   address,
			CreatedAt: now,
		}

		mockGateway.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
				assert.Equal(t, address, wallet.Address)
				return expectedWallet, nil
			})

		output, err := interactor.Get(context.Background(), &ports.WalletInput{Address: address})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, id, output.ID)
		assert.Equal(t, address, output.Address)
	})

	t.Run("異常系: DBエラー時にエラーが返る", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockGateway := mock.NewMockWalletGateway(ctrl)
		interactor := NewWalletInteractor(mockGateway)

		mockGateway.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("db error"))

		output, err := interactor.Get(context.Background(), &ports.WalletInput{Address: address})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "db error", err.Error())
	})
}

func TestWalletInteractor_Create(t *testing.T) {
	address := "0x29BE6Cd779c7ddbF6cBe3efdeF9Ba6d9CD2C7W4D"

	t.Run("正常系: ウォレットを作成できる", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockGateway := mock.NewMockWalletGateway(ctrl)
		interactor := NewWalletInteractor(mockGateway)

		mockGateway.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil)

		output, err := interactor.Create(context.Background(), &ports.WalletInput{Address: address})

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, address, output.Address)
	})

	t.Run("異常系: Gateway.Createでエラーが発生した場合", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockGateway := mock.NewMockWalletGateway(ctrl)
		interactor := NewWalletInteractor(mockGateway)

		mockGateway.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("create error"))

		output, err := interactor.Create(context.Background(), &ports.WalletInput{Address: address})

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "create error", err.Error())
	})
}
