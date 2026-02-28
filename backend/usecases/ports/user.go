// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import (
	"time"

	"github.com/google/uuid"
)

// UserInput はコントローラから取得する構造体を表します。
type UserInput struct {
	Name       string    `json:"name" validate:"required" example:"北川武"`
	Email      string    `json:"email" validate:"required,email" example:"takeshi-kitagawa@gmail.com"`
	Wallet     string    `json:"wallet" validate:"required" example:"0x411F8Cd035DB4874a8F1247Cb49908c794A0AB50"`
	Address    string    `json:"address" example:"東京都千代田区１−１−１"`
	BusinessID uuid.UUID `json:"business_id" example:"019504de-f120-76af-905b-2fe06352b04a"`
	Website    string    `json:"website" example:"https://www.yahoo.com"`
	FaceImage  string    `json:"face_image" example:"https://www.yahoo.com/img/test.jpg"`
	Eyecatch   string    `json:"eyecatch" example:"https://www.yahoo.com/img/test.jpg"`
	Profile    string    `json:"profile" validate:"min=10,max=4000" example:"私はいつでも明るいです"`
	Role       string    `json:"role" example:"member"`
}

// UserOutput はAPIで返す構造体
type UserOutput struct {
	ID         uuid.UUID `json:"id" example:"01932563-f671-71ff-9a0d-c452de9d06aa"`
	Name       string    `json:"name" example:"北川武"`
	Email      string    `json:"email" example:"takeshi-kitagawa@gmail.com"`
	Wallet     string    `json:"wallet" example:"01932564-3432-7b56-b772-7c69586e1897"`
	Address    string    `json:"address,omitempty" example:"東京都千代田区１−１−１"`
	BusinessID uuid.UUID `json:"business_id" example:"019504de-f120-76af-905b-2fe06352b04a"`
	Website    string    `json:"website,omitempty" example:"https://www.yahoo.com"`
	FaceImage  string    `json:"face_image,omitempty" example:"https://www.yahoo.com/img/test.jpg"`
	Eyecatch   string    `json:"eyecatch,omitempty" example:"https://www.yahoo.com/img/test.jpg"`
	Profile    string    `json:"profile,omitempty" example:"私はいつでも明るいです"`
	Role       string    `json:"role" example:"member"`
	CreatedAt  time.Time `json:"created_at" example:"2024-11-04T20:51:26Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2024-11-04T20:51:26Z"`
}
