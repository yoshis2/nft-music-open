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

func TestBusinessInteractor(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Run("正常系: ビジネス情報を追加できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			input := &ports.BusinessMasterInput{Name: "Composer"}

			mockGateway.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(nil)

			output, err := interactor.Create(context.Background(), input)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, "Composer", output.Name)
		})

		t.Run("異常系: Gateway.Createでエラーが発生した場合", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			input := &ports.BusinessMasterInput{Name: "Composer"}

			mockGateway.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(errors.New("db error"))

			output, err := interactor.Create(context.Background(), input)

			assert.Error(t, err)
			assert.Nil(t, output)
			assert.Equal(t, "db error", err.Error())
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("正常系: IDでビジネス情報を取得できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			id := uuid.New()
			expectedDomain := &domain.BusinessMaster{
				ID:        id,
				Name:      "Vocalist",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			mockGateway.EXPECT().
				Get(gomock.Any(), id).
				Return(expectedDomain, nil)

			output, err := interactor.Get(context.Background(), id)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, id, output.ID)
			assert.Equal(t, "Vocalist", output.Name)
		})

		t.Run("異常系: Gateway.Getでエラーが発生した場合", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			id := uuid.New()
			mockGateway.EXPECT().
				Get(gomock.Any(), id).
				Return(nil, errors.New("not found"))

			output, err := interactor.Get(context.Background(), id)

			assert.Error(t, err)
			assert.Nil(t, output)
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("正常系: 一覧を取得できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			expectedDomains := []domain.BusinessMaster{
				{ID: uuid.New(), Name: "A"},
				{ID: uuid.New(), Name: "B"},
			}

			mockGateway.EXPECT().
				List(gomock.Any()).
				Return(expectedDomains, nil)

			outputs, err := interactor.List(context.Background())

			assert.NoError(t, err)
			assert.Equal(t, 2, len(outputs))
			assert.Equal(t, "A", outputs[0].Name)
			assert.Equal(t, "B", outputs[1].Name)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("正常系: 情報を更新できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			id := uuid.New()
			input := &ports.BusinessMasterInput{Name: "New Name"}

			mockGateway.EXPECT().
				Update(gomock.Any(), gomock.Any()).
				Return(nil)

			output, err := interactor.Update(context.Background(), id, input)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, "New Name", output.Name)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("正常系: 削除できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockBusinessGateway(ctrl)
			interactor := NewBusinessInteractor(mockGateway)

			id := uuid.New()

			mockGateway.EXPECT().
				Delete(gomock.Any(), gomock.Any()).
				Return(nil)

			err := interactor.Delete(context.Background(), id)

			assert.NoError(t, err)
		})
	})
}
