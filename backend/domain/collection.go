package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID              uuid.UUID      `gorm:"id"`
	UserID          uuid.UUID      `gorm:"user_id"`
	ChainID         int            `gorm:"chain_id"`
	Name            string         `gorm:"name"`
	ContractAddress string         `gorm:"contract_address"`
	Description     sql.NullString `gorm:"description"`
	ImageURL        sql.NullString `gorm:"image_url" swaggertype:"string"`
	BannerImageURL  sql.NullString `gorm:"banner_image_url" swaggertype:"string"`
	ExternalURL     sql.NullString `gorm:"external_url" swaggertype:"string"`
	Royalty         int            `gorm:"royalty"`
	RoyaltyReceiver string         `gorm:"royalty_receiver" swaggertype:"string"`
	CreatedAt       time.Time      `gorm:"created_at"`
	UpdatedAt       time.Time      `gorm:"updated_at"`
}
