// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"nft-music/infrastructure/logging"
	"nft-music/usecases/ports"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWalletPresenter_Exist(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/wallet", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	logging := logging.NewZapLogging()

	presenter := NewWalletPresenter(logging)
	id1, _ := uuid.Parse("0192fe69-3de0-7465-ab55-ad1476f83931")

	output := &ports.WalletOutput{ID: id1}

	if assert.NoError(t, presenter.Exist(c, output)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "ウォレットは既に存在しています")
	}
}

func TestWalletPresenter_Create(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/wallet", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	logging := logging.NewZapLogging()
	presenter := NewWalletPresenter(logging)

	id := "0192fe69-3de0-7465-ab55-ad1476f83932"
	idUUID, _ := uuid.Parse(id)
	output := &ports.WalletOutput{ID: idUUID}

	if assert.NoError(t, presenter.Create(c, output)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), id)
	}
}
