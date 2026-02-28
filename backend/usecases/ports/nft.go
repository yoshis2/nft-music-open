// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

// NftInput はコントローラから取得する構造体を表します。
type NftInput struct {
	ChainID     int       `json:"chain_id" validate:"required" example:"222"`
	Wallet      string    `json:"wallet" validate:"required" example:"0xc5309Ef694C81C4a8e946F2810e09516436daeB5"`
	Name        string    `json:"name" validate:"required" example:"GoodNFT"`
	Description string    `json:"description" validate:"required" example:"良いNFTです"`
	FileType    string    `json:"file_type" example:"audio"`
	ImageCid    string    `json:"image_cid" validate:"omitempty" example:"QmW9qWKbZneDijFPcHdbn41fQiHRwownM4U6avLsCHfMJS"`
	AudioCid    string    `json:"audio_cid" validate:"omitempty" example:"QmW9qWKbZneDijFPcHdbn41fQiHRwownM4U6avLsCHfMJS"`
	VideoCid    string    `json:"video_cid" validate:"omitempty" example:"QmW9qWKbZneDijFPcHdbn41fQiHRwownM4U6avLsCHfMJS"`
	GenreID     uuid.UUID `json:"genre_id" validate:"required" example:"019504e3-d996-7979-8043-ef03fa7a6d89"`
	Status      string    `json:"status" validate:"required" example:"mint"`
	Price       float64   `json:"price,string" validate:"required" example:"1000.11"`
	Insentive   int       `json:"insentive,string" validate:"required" example:"20"`
	Sale        bool      `json:"sale" example:"0"`
}

// NftOutput はAPIで返す構造体
type NftOutput struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	ChainID       int       `json:"chain_id"`
	FileType      string    `json:"file_type"`
	TransactionID string    `json:"transaction_id"`
	TokenURL      string    `json:"token_url"`
	GenreID       uuid.UUID `json:"genre_id"`
	Status        string    `json:"status"`
	Price         float64   `json:"price"`
	Insentive     int       `json:"insentive"`
	Sale          bool      `json:"sale"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
