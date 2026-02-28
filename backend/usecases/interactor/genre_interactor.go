// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"

	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/google/uuid"
)

// GenreInteractor ジャンルインストラクタの構造体
type GenreInteractor struct {
	Gateway gateways.GenreGateway
}

func NewGenreInteractor(gateway gateways.GenreGateway) *GenreInteractor {
	return &GenreInteractor{
		Gateway: gateway,
	}
}

// Create はジャンルの追加をする
func (interactor *GenreInteractor) Create(ctx context.Context, input *ports.GenreMasterInput) (*ports.GenreMasterOutput, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := util.JapaneseNowTime()

	genre := &domain.GenreMaster{
		ID:        uuidV7,
		Name:      input.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := interactor.Gateway.Create(ctx, genre); err != nil {
		return nil, err
	}

	output := &ports.GenreMasterOutput{
		ID:        genre.ID,
		Name:      genre.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}

	return output, nil
}

// Get は指定のジャンルを一つ取得する
func (interactor *GenreInteractor) Get(ctx context.Context, id uuid.UUID) (*ports.GenreMasterOutput, error) {
	domainOutput, err := interactor.Gateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &ports.GenreMasterOutput{
		ID:        domainOutput.ID,
		Name:      domainOutput.Name,
		CreatedAt: domainOutput.CreatedAt,
		UpdatedAt: domainOutput.UpdatedAt,
	}
	return output, nil
}

// List はジャンルのマスター情報を一覧で取得する
func (interactor *GenreInteractor) List(ctx context.Context) ([]ports.GenreMasterOutput, error) {
	genres, err := interactor.Gateway.List(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []ports.GenreMasterOutput
	for _, genre := range genres {
		output := ports.GenreMasterOutput{
			ID:        genre.ID,
			Name:      genre.Name,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
		}
		outputs = append(outputs, output)
	}

	return outputs, nil
}

// Update はジャンルの情報を変更する
func (interactor *GenreInteractor) Update(ctx context.Context, id uuid.UUID, input *ports.GenreMasterInput) (*ports.GenreMasterOutput, error) {
	now := util.JapaneseNowTime()
	genre := &domain.GenreMaster{
		ID:        id,
		Name:      input.Name,
		UpdatedAt: now,
	}

	if err := interactor.Gateway.Update(ctx, genre); err != nil {
		return nil, err
	}

	output := &ports.GenreMasterOutput{
		ID:        genre.ID,
		Name:      genre.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}

	return output, nil
}

// Delete は指定した一つのジャンルを削除する
func (interactor *GenreInteractor) Delete(ctx context.Context, id uuid.UUID) error {
	genre := domain.GenreMaster{
		ID: id,
	}
	return interactor.Gateway.Delete(ctx, &genre)
}
