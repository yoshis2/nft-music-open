// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

type CollectionInput struct {
	UserID          uuid.UUID `json:"user_id" example:"0193254c-a151-7c4c-b06a-259da7258a27"`
	ChainID         int       `json:"chain_id" example:"1"`
	Name            string    `json:"name" example:"建築"`
	ContractAddress string    `json:"contract_address" example:"0x495f947276749ce646f68ac8c248420045075b34"`
	Description     string    `json:"description" example:"私はいつでも明るいです"`
	ImageURL        string    `json:"image_url" example:"https://www.yahoo.com/img/test.jpg"`
	BannerImageURL  string    `json:"banner_image_url" example:"https://www.yahoo.com/img/test.jpg"`
	ExternalURL     string    `json:"external_url" example:"https://www.yahoo.com"`
	Royalty         int       `json:"royalty" example:"10"`
	RoyaltyReceiver string    `json:"royalty_receiver" example:"0x495f947276749ce646f68ac8c248420045075b34"`
}

type CollectionOutput struct {
	ID              uuid.UUID `json:"id" example:"0193254c-a151-7c4c-b06a-259da7258a27"`
	UserID          uuid.UUID `json:"user_id" example:"0193254c-a151-7c4c-b06a-259da7258a27"`
	ChainID         int       `json:"chain_id" example:"1"`
	Name            string    `json:"name" example:"建築"`
	ContractAddress string    `json:"contract_address" example:"0x495f947276749ce646f68ac8c248420045075b34"`
	Description     string    `json:"description" example:"私はいつでも明るいです"`
	ImageURL        string    `json:"image_url" example:"https://www.yahoo.com/img/test.jpg"`
	BannerImageURL  string    `json:"banner_image_url" example:"https://www.yahoo.com/img/test.jpg"`
	ExternalURL     string    `json:"external_url" example:"https://www.yahoo.com"`
	Royalty         int       `json:"royalty" example:"10"`
	RoyaltyReceiver string    `json:"royalty_receiver" example:"0x495f947276749ce646f68ac8c248420045075b34"`
	CreatedAt       time.Time `json:"created_at" example:"2024-11-04T20:51:26Z"`
	UpdatedAt       time.Time `json:"updated_at" example:"2024-11-04T20:51:26Z"`
}
