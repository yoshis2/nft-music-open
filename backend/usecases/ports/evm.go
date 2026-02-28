// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

type EVMInput struct{}

type SignerOutput struct {
	AddressHex      string `json:"address_hex" example:"0x47CD2D0873833Ba015e0C31AB94D30313eF07942"`
	TransactionHash string `json:"transaction_hash" example:"0x9d605694113bde48dd07bf4d6f408906d3d5a9ee6656df600209ad6a38b2669e"`
}
