// Package ports は、ユースケースで使用される入力・出力モデルを定義します。
package ports

import "github.com/google/uuid"

// CreatedObject 既に存在する時に返却するオブジェクト
type CreatedObject struct {
	ID         uuid.UUID `json:"id" example:"01932564-3432-7b56-b772-7c69586e1897"`
	StatusCode int       `json:"status_code" example:"201"`
	Message    string    `json:"message" example:"データベースには既に存在しています"`
}

// ErrorResponseObject error時に返却するオブジェクト
type ErrorResponseObject struct {
	StatusCode int    `json:"status_code" example:"0"`
	ErrorType  string `json:"error_type" example:"リクエストの構文が不正やページが見つからない等のタイプ別"`
	Message    string `json:"message" example:"原因となるのエラーメッセージ"`
}
