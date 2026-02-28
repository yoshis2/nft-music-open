// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"net/http"

	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/labstack/echo/v4"
)

// WalletPresenter の構造体
type WalletPresenter struct {
	Logging logging.Logging
}

// NewWalletPresenter WalletPresenterのコンストラクタ
func NewWalletPresenter(logging logging.Logging) *WalletPresenter {
	return &WalletPresenter{
		Logging: logging,
	}
}

// Exist はもし、ウォレットアドレスが既に登録されていたら作成済みを返す
func (presenter *WalletPresenter) Exist(c echo.Context, output *ports.WalletOutput) error {
	presenter.Logging.Info("ウォレットは既に存在しています")
	created := ports.CreatedObject{
		ID:         output.ID,
		StatusCode: http.StatusCreated,
		Message:    "ウォレットは既に存在しています",
	}
	return c.JSON(http.StatusCreated, &created)
}

// Create はウォレットアドレスの登録をWebに返す
func (presenter *WalletPresenter) Create(c echo.Context, output *ports.WalletOutput) error {
	presenter.Logging.Info("ウォレットを登録しました")
	return c.JSON(http.StatusOK, output)
}
