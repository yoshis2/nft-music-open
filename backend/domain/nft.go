package domain

import (
	"time"

	"github.com/google/uuid"
)

// NFT はNFT構造体
type NFT struct {
	ID            uuid.UUID `gorm:"id"`
	UserID        uuid.UUID `gorm:"user_id"`
	ChainID       string    `gorm:"chain_id"`
	TransactionID string    `gorm:"transaction_id"`
	MetajsonURL   string    `gorm:"metajson_url"`
	GenreID       uuid.UUID `gorm:"genre_id"`
	Status        string    `gorm:"status"`
	Sale          bool      `gorm:"sale"`
	Price         float64   `gorm:"price"`
	Insentive     int       `gorm:"insentive"`
	CreatedAt     time.Time `gorm:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at"`
}
