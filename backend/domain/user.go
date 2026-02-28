// Package domain は、ドメインモデルとビジネスルールを定義します。
package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// User はユーザーの構造体です
type User struct {
	ID         uuid.UUID      `gorm:"id"`
	Wallet     string         `gorm:"wallet"`
	Name       string         `gorm:"name"`
	Email      string         `gorm:"email"`
	Address    sql.NullString `gorm:"address"`
	BusinessID uuid.UUID      `gorm:"business_id"`
	Website    sql.NullString `gorm:"website"`
	FaceImage  sql.NullString `gorm:"face_image"`
	Eyecatch   sql.NullString `gorm:"eyecatch"`
	Profile    sql.NullString `gorm:"profile"`
	Role       string         `gorm:"role"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdatedAt  time.Time      `gorm:"updated_at"`
}
