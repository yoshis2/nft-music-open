// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

// WalletInput はコントローラから取得する構造体を表します。
type WalletInput struct {
	Address string `json:"address" validate:"required" example:"0x29BE6Cd779c7ddbF6cBe3efdeF7Ba6f9CD2C703D"`
}

// WalletOutput はAPIで返す構造体
type WalletOutput struct {
	ID        uuid.UUID `json:"id" example:"01932550-9ecc-774f-b1ec-98d3002b9b2e"`
	Address   string    `json:"address" example:"0x85A4599940764d07f469EdEdAb10a72B727B6669"`
	CreatedAt time.Time `json:"created_at" example:"2024-11-07T21:51:36Z"`
}
