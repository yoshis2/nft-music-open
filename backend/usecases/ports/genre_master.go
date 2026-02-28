// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

// GenreMasterInput はコントローラから取得する構造体を表します。
type GenreMasterInput struct {
	Name string `json:"name" validate:"required" example:"エレクトロニック・ダンス・ミュージック"`
}

// GenreMasterOutput はAPIで返す構造体
type GenreMasterOutput struct {
	ID        uuid.UUID `json:"id" example:"01932573-3f06-78c6-9743-b495cfcc2167"`
	Name      string    `json:"name" example:"エレクトロニック・ダンス・ミュージック"`
	CreatedAt time.Time `json:"created_at" example:"2024-11-04T20:51:26Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-11-04T20:51:26Z"`
}
