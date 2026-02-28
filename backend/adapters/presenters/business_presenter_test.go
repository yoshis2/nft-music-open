// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"nft-music/infrastructure/logging"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBusinessPresenter_Create(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/businesses", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	logging := logging.NewZapLogging()

	presenter := NewBusinessPresenter(logging)

	if assert.NoError(t, presenter.Create(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
