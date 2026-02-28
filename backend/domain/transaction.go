package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              string         `gorm:"id"`
	UserID          uuid.UUID      `gorm:"user_id"`
	ChainID         int            `gorm:"chain_id"`
	ContractAddress string         `gorm:"contract_address"`
	Nonce           int            `gorm:"nonce"`
	TokenURL        string         `gorm:"token_url"`
	GenreID         uuid.UUID      `gorm:"genre_id"`
	To              sql.NullString `gorm:"to"`
	Price           float64        `gorm:"price"` // Value
	Insentive       int            `gorm:"insentive"`
	Cost            int            `gorm:"cost"`
	Sale            bool           `gorm:"sale"`
	Status          string         `gorm:"status"`
	CreatorAddress  string         `gorm:"creator_address"`
	CreatedAt       time.Time      `gorm:"created_at"`
	UpdatedAt       time.Time      `gorm:"updated_at"`
}
