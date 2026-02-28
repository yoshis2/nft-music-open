// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"database/sql"

	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/google/uuid"
)

// CollectionInteractor コレクションインストラクタの構造体
type CollectionInteractor struct {
	Gateway gateways.CollectionGateway
}

func NewCollectionInteractor(gateway gateways.CollectionGateway) *CollectionInteractor {
	return &CollectionInteractor{
		Gateway: gateway,
	}
}

func (interactor *CollectionInteractor) Get(ctx context.Context, id uuid.UUID) (*ports.CollectionOutput, error) {
	collection, err := interactor.Gateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return output(collection), nil
}

func (interactor *CollectionInteractor) List(ctx context.Context) ([]ports.CollectionOutput, error) {
	collections, err := interactor.Gateway.List(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []ports.CollectionOutput
	for _, collection := range collections {
		outputs = append(outputs, *output(&collection))
	}
	return outputs, nil
}

// Create コレクションを作成する
func (interactor *CollectionInteractor) Create(ctx context.Context, input *ports.CollectionInput) (*ports.CollectionOutput, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	collection := &domain.Collection{
		ID:              uuidV7,
		UserID:          input.UserID,
		ChainID:         input.ChainID,
		Name:            input.Name,
		ContractAddress: input.ContractAddress,
		Description:     sql.NullString{String: input.Description, Valid: true},
		ImageURL:        sql.NullString{String: input.ImageURL, Valid: true},
		BannerImageURL:  sql.NullString{String: input.BannerImageURL, Valid: true},
		ExternalURL:     sql.NullString{String: input.ExternalURL, Valid: true},
		Royalty:         input.Royalty,
		RoyaltyReceiver: input.RoyaltyReceiver,
		CreatedAt:       util.JapaneseNowTime(),
		UpdatedAt:       util.JapaneseNowTime(),
	}

	if err := interactor.Gateway.Create(ctx, collection); err != nil {
		return nil, err
	}

	return output(collection), nil
}

func (interactor *CollectionInteractor) Update(ctx context.Context, id uuid.UUID, input *ports.CollectionInput) (*ports.CollectionOutput, error) {
	now := util.JapaneseNowTime()
	collection := &domain.Collection{
		ID:              id,
		UserID:          input.UserID,
		ChainID:         input.ChainID,
		Name:            input.Name,
		ContractAddress: input.ContractAddress,
		Description:     sql.NullString{String: input.Description, Valid: true},
		ImageURL:        sql.NullString{String: input.ImageURL, Valid: true},
		BannerImageURL:  sql.NullString{String: input.BannerImageURL, Valid: true},
		ExternalURL:     sql.NullString{String: input.ExternalURL, Valid: true},
		Royalty:         input.Royalty,
		RoyaltyReceiver: input.RoyaltyReceiver,
		UpdatedAt:       now,
	}

	if err := interactor.Gateway.Update(ctx, collection); err != nil {
		return nil, err
	}

	return output(collection), nil
}

func (interactor *CollectionInteractor) Delete(ctx context.Context, id uuid.UUID) error {
	collection := &domain.Collection{
		ID: id,
	}
	return interactor.Gateway.Delete(ctx, collection)
}

func output(collection *domain.Collection) *ports.CollectionOutput {
	return &ports.CollectionOutput{
		ID:              collection.ID,
		UserID:          collection.UserID,
		ChainID:         collection.ChainID,
		Name:            collection.Name,
		ContractAddress: collection.ContractAddress,
		Description:     collection.Description.String,
		ImageURL:        collection.ImageURL.String,
		BannerImageURL:  collection.BannerImageURL.String,
		ExternalURL:     collection.ExternalURL.String,
		Royalty:         collection.Royalty,
		RoyaltyReceiver: collection.RoyaltyReceiver,
		CreatedAt:       collection.CreatedAt,
		UpdatedAt:       collection.UpdatedAt,
	}
}
