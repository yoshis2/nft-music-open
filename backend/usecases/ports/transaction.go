// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

type TransactionOutput struct {
	ID          string    `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ChainID     int       `json:"chain_id"`
	TokenID     int       `json:"token_id"`
	Nonce       int       `json:"nonce"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FileType    string    `json:"file_type"`
	ImageURL    string    `json:"image_url"`
	AudioURL    string    `json:"audio_url"`
	VideoURL    string    `json:"video_url"`
	TokenURL    string    `json:"token_url"`
	GenreID     uuid.UUID `json:"genre_id"`
	GenreName   string    `json:"genre_name"`
	To          string    `json:"to"`
	Price       float64   `json:"price"`
	Insentive   int       `json:"insentive"`
	Cost        int       `json:"cost"`
	Sale        bool      `json:"sale"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
