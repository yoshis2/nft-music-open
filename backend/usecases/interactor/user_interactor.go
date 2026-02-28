// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"database/sql"

	"nft-music/adapters/presenters"
	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/google/uuid"
)

type UserInteractor struct {
	UserGateway gateways.UserGateway
	Logging     logging.Logging
	Error       *presenters.ErrorPresenter
}

func NewUserInteractor(userGateway gateways.UserGateway, logging logging.Logging) *UserInteractor {
	return &UserInteractor{
		UserGateway: userGateway,
		Logging:     logging,
		Error:       presenters.NewErrorPresenter(logging),
	}
}

func (interactor *UserInteractor) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &domain.User{
		ID: id,
	}
	return interactor.UserGateway.Get(ctx, user)
}

func (interactor *UserInteractor) GetByWallet(ctx context.Context, wallet string) (*domain.User, error) {
	user := &domain.User{
		Wallet: wallet,
	}
	return interactor.UserGateway.GetByWallet(ctx, user)
}

func (interactor *UserInteractor) List(ctx context.Context) ([]domain.User, error) {
	return interactor.UserGateway.List(ctx)
}

func (interactor *UserInteractor) Create(ctx context.Context, input ports.UserInput) error {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil
	}
	now := util.JapaneseNowTime()
	user := &domain.User{
		ID:     uuidV7,
		Name:   input.Name,
		Email:  input.Email,
		Wallet: input.Wallet,
		Address: sql.NullString{
			String: input.Address,
			Valid:  true,
		},
		BusinessID: input.BusinessID,
		Website: sql.NullString{
			String: input.Website,
			Valid:  true,
		},
		FaceImage: sql.NullString{
			String: input.FaceImage,
			Valid:  true,
		},
		Eyecatch: sql.NullString{
			String: input.Eyecatch,
			Valid:  true,
		},
		Profile: sql.NullString{
			String: input.Profile,
			Valid:  true,
		},
		Role:      "member",
		CreatedAt: now,
		UpdatedAt: now,
	}

	return interactor.UserGateway.Create(ctx, user)
}

func (interactor *UserInteractor) Update(ctx context.Context, id uuid.UUID, input ports.UserInput) error {
	user := &domain.User{
		ID:     id,
		Name:   input.Name,
		Email:  input.Email,
		Wallet: input.Wallet,
		Address: sql.NullString{
			String: input.Address,
			Valid:  true,
		},
		BusinessID: input.BusinessID,
		Website: sql.NullString{
			String: input.Website,
			Valid:  true,
		},
		FaceImage: sql.NullString{
			String: input.FaceImage,
			Valid:  true,
		},
		Eyecatch: sql.NullString{
			String: input.Eyecatch,
			Valid:  true,
		},
		Profile: sql.NullString{
			String: input.Profile,
			Valid:  true,
		},
		Role:      "member",
		UpdatedAt: util.JapaneseNowTime(),
	}

	return interactor.UserGateway.Update(ctx, user)
}

func (interactor *UserInteractor) Delete(ctx context.Context, id uuid.UUID) error {
	user := &domain.User{
		ID: id,
	}
	return interactor.UserGateway.Delete(ctx, user)
}
