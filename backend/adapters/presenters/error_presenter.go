// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"net/http"
	"strings"

	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/labstack/echo/v4"
)

// ErrorPresenter の構造体
type ErrorPresenter struct {
	Logging logging.Logging
}

// NewErrorPresenter のコンストラクタ
func NewErrorPresenter(logging logging.Logging) *ErrorPresenter {
	return &ErrorPresenter{
		Logging: logging,
	}
}

// ErrorResponse はエラーレスポンス
func (presenter *ErrorPresenter) ErrorResponse(c echo.Context, err error) error {
	var code int
	var errorType string
	if err != nil {
		if isBadRequestError(err.Error()) {
			code = http.StatusBadRequest
			errorType = "リクエストの構文が不正 : " + err.Error()
		} else if isUnauthorizedError(err.Error()) {
			code = http.StatusUnauthorized
			errorType = "保護された API への認証されていないリクエスト"
		} else if isRecordNotFoundError(err.Error()) {
			code = http.StatusNotFound
			errorType = "DBのテーブルにデータがありません"
		} else if isNotFoundError(err.Error()) {
			code = http.StatusNotFound
			errorType = "ページ(APIにデータ)が見つからない"
		} else if isDuplicatedUError(err.Error()) {
			code = http.StatusConflict
			errorType = "現在のサーバーの状態と競合"
		} else if isCreated(err.Error()) {
			code = http.StatusCreated
			errorType = "正常"
		} else {
			code = http.StatusInternalServerError
			errorType = "サーバーが予期せぬ状況に遭遇し、リクエストが履行されなかった"
		}
	}

	presenter.Logging.Error("Message: " + err.Error())
	return c.JSON(code, ports.ErrorResponseObject{
		StatusCode: code,
		ErrorType:  errorType,
		Message:    err.Error(),
	})
}

func isBadRequestError(msg string) bool {
	var messageBool bool

	if strings.Contains(msg, "a foreign key constraint fails") {
		messageBool = true
	} else if strings.Contains(msg, "Error:Field validation") {
		messageBool = true
	} else if strings.Contains(msg, "Out of range value for") {
		messageBool = true
	} else if strings.Contains(msg, "BadRequest") {
		messageBool = true
	}
	return messageBool
}

func isRecordNotFoundError(msg string) bool {
	return strings.Contains(msg, "record not found")
}

func isNotFoundError(msg string) bool {
	var messageBool bool

	if strings.Contains(msg, "Not Found") {
		messageBool = true
	} else if strings.Contains(msg, "doesn't exist") {
		messageBool = true
	}

	return messageBool
}

func isDuplicatedUError(msg string) bool {
	return strings.Contains(msg, "Already Exist")
}

func isUnauthorizedError(msg string) bool {
	return strings.Contains(msg, "Unauthorized")
}

func isCreated(msg string) bool {
	return strings.Contains(msg, "Already Created")
}
