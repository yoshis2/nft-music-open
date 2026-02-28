// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

// BusinessMasterInput はコントローラから取得する構造体を表します。
type BusinessMasterInput struct {
	Name string `json:"name" validate:"required" example:"建築"`
}

// BusinessMasterOutput はAPIで返す構造体
type BusinessMasterOutput struct {
	ID        uuid.UUID `json:"id" example:"0193254c-a151-7c4c-b06a-259da7258a27"`
	Name      string    `json:"name" example:"建築"`
	CreatedAt time.Time `json:"created_at" example:"2024-11-04T20:51:26Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-11-04T20:51:26Z"`
}
