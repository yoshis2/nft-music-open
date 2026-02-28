// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"nft-music/domain"
	"nft-music/usecases/gateways/mock"
	"nft-music/usecases/ports"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCollectionInteractor(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		t.Run("正常系: IDでコレクションを取得できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

			id := uuid.New()
			expectedDomain := &domain.Collection{
				ID:          id,
				Name:        "Test Collection",
				Description: sql.NullString{String: "Desc", Valid: true},
			}

			mockGateway.EXPECT().
				Get(gomock.Any(), id).
				Return(expectedDomain, nil)

			output, err := interactor.Get(context.Background(), id)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, id, output.ID)
			assert.Equal(t, "Test Collection", output.Name)
			assert.Equal(t, "Desc", output.Description)
		})

		t.Run("異常系: Gateway.Getでエラーが発生した場合", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

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
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

			expectedDomains := []domain.Collection{
				{ID: uuid.New(), Name: "Col 1"},
				{ID: uuid.New(), Name: "Col 2"},
			}

			mockGateway.EXPECT().
				List(gomock.Any()).
				Return(expectedDomains, nil)

			outputs, err := interactor.List(context.Background())

			assert.NoError(t, err)
			assert.Equal(t, 2, len(outputs))
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("正常系: コレクションを作成できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

			input := &ports.CollectionInput{
				Name: "New Collection",
			}

			mockGateway.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(nil)

			output, err := interactor.Create(context.Background(), input)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, "New Collection", output.Name)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("正常系: コレクションを更新できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

			id := uuid.New()
			input := &ports.CollectionInput{
				Name: "Updated Name",
			}

			mockGateway.EXPECT().
				Update(gomock.Any(), gomock.Any()).
				Return(nil)

			output, err := interactor.Update(context.Background(), id, input)

			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, "Updated Name", output.Name)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("正常系: 削除できる", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockGateway := mock.NewMockCollectionGateway(ctrl)
			interactor := NewCollectionInteractor(mockGateway)

			id := uuid.New()

			mockGateway.EXPECT().
				Delete(gomock.Any(), gomock.Any()).
				Return(nil)

			err := interactor.Delete(context.Background(), id)

			assert.NoError(t, err)
		})
	})
}
