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

// BusinessInteractor ビジネスインストラクタの構造体
type BusinessInteractor struct {
	Gateway gateways.BusinessGateway
}

// NewBusinessInteractor はBusinessInteractorを初期化します。
func NewBusinessInteractor(gateway gateways.BusinessGateway) *BusinessInteractor {
	return &BusinessInteractor{
		Gateway: gateway,
	}
}

// Create は職種マスターを追加する
func (interactor *BusinessInteractor) Create(ctx context.Context, input *ports.BusinessMasterInput) (*ports.BusinessMasterOutput, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := util.JapaneseNowTime()

	business := &domain.BusinessMaster{
		ID:        uuidV7,
		Name:      input.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := interactor.Gateway.Create(ctx, business); err != nil {
		return nil, err
	}

	output := &ports.BusinessMasterOutput{
		ID:        business.ID,
		Name:      business.Name,
		CreatedAt: business.CreatedAt,
		UpdatedAt: business.UpdatedAt,
	}

	return output, nil
}

// Get は職種マスターの存在を確認する
func (interactor *BusinessInteractor) Get(ctx context.Context, id uuid.UUID) (*ports.BusinessMasterOutput, error) {
	domainOutput, err := interactor.Gateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &ports.BusinessMasterOutput{
		ID:        domainOutput.ID,
		Name:      domainOutput.Name,
		CreatedAt: domainOutput.CreatedAt,
		UpdatedAt: domainOutput.UpdatedAt,
	}
	return output, nil
}

// List は職種マスターの一覧を取得
func (interactor *BusinessInteractor) List(ctx context.Context) ([]ports.BusinessMasterOutput, error) {
	businesses, err := interactor.Gateway.List(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []ports.BusinessMasterOutput
	for _, business := range businesses {
		output := ports.BusinessMasterOutput{
			ID:        business.ID,
			Name:      business.Name,
			CreatedAt: business.CreatedAt,
			UpdatedAt: business.UpdatedAt,
		}
		outputs = append(outputs, output)
	}

	return outputs, nil
}

// Update は職種マスターを修正する
func (interactor *BusinessInteractor) Update(ctx context.Context, id uuid.UUID, input *ports.BusinessMasterInput) (*ports.BusinessMasterOutput, error) {
	now := util.JapaneseNowTime()
	business := &domain.BusinessMaster{
		ID:        id,
		Name:      input.Name,
		UpdatedAt: now,
	}

	if err := interactor.Gateway.Update(ctx, business); err != nil {
		return nil, err
	}

	output := &ports.BusinessMasterOutput{
		ID:        business.ID,
		Name:      business.Name,
		CreatedAt: business.CreatedAt,
		UpdatedAt: business.UpdatedAt,
	}

	return output, nil
}

// Delete は職種マスターを削除する
func (interactor *BusinessInteractor) Delete(ctx context.Context, id uuid.UUID) error {
	business := domain.BusinessMaster{
		ID: id,
	}
	return interactor.Gateway.Delete(ctx, &business)
}
