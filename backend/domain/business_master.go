package domain

import (
	"time"

	"github.com/google/uuid"
)

// BusinessMaster は職種マスターの構造体
type BusinessMaster struct {
	ID        uuid.UUID `gorm:"id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
