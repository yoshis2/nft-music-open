package domain

import (
	"time"

	"github.com/google/uuid"
)

// GenreMaster は音楽ジャンルマスターの構造体
type GenreMaster struct {
	ID        uuid.UUID `gorm:"id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
