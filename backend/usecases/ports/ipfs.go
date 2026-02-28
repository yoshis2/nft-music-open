// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import "github.com/google/uuid"

// IpfsInput はコントローラーから取得する構造体
type IpfsInput struct {
	Wallet string `form:"wallet"`
	File   string `form:"file"`
}

type IpfsMetaInput struct {
	Name        string `json:"name" validate:"required" example:"GoodNFT"`
	Description string `json:"description" validate:"required" example:"良いNFTです"`
	FileType    string `json:"file_type" validate:"required" example:"audio"`
	ImageCid    string `json:"image_cid" validate:"required" example:"QmW9qWKbZneDijFPcHdbn41fQiHRwownM4U6avLsCHfMJS"`
	AudioCid    string `json:"audio_cid" validate:"required" example:"QmPYVinZ3fmWSetwc9FWMF9b93hfKjWiHtygXfVY8exwEM"`
	VideoCid    string `json:"video_cid" validate:"required" example:"QmQfESxcop5u9zjvK5hJyX8oKJJZCsEYVh4kTU64Wzmxqz"`
	Insentive   int    `json:"insentive" example:"10"`
}

// IpfsOutput はコントローラへ返す構造体
type IpfsOutput struct {
	UserID uuid.UUID `json:"user_id"`
	Cid    string    `json:"cid"`
	Path   string    `json:"path"`
}
