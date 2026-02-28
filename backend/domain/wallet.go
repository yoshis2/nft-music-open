package domain

import (
	"time"

	"github.com/google/uuid"
)

// Wallet はウォレットアドレスの構造体
type Wallet struct {
	ID        uuid.UUID `gorm:"id"`
	Address   string    `gorm:"address"`
	CreatedAt time.Time `gorm:"created_at"`
}
