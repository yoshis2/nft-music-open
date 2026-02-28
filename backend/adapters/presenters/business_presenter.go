// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"net/http"

	"nft-music/usecases/logging"

	"github.com/labstack/echo/v4"
)

// BusinessPresenter は表示用の構造体
type BusinessPresenter struct {
	Logging logging.Logging
}

// NewBusinessPresenter BusinessPresenterのコンストラクタ
func NewBusinessPresenter(logging logging.Logging) *BusinessPresenter {
	return &BusinessPresenter{
		Logging: logging,
	}
}

// Create は職種の登録をWebに返す
func (presenter *BusinessPresenter) Create(c echo.Context) error {
	presenter.Logging.Info("職種を登録しました")
	return c.JSON(http.StatusOK, nil)
}
