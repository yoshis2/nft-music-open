// Package presenters は、APIレスポンスのデータ整形を行います。
package presenters

import (
	"testing"

	"nft-music/infrastructure/logging"
)

// TestNewErrorPresenter のテスト
func TestNewErrorPresenter(t *testing.T) {
	logging := logging.NewZapLogging()
	presenter := NewErrorPresenter(logging)

	if presenter.Logging == nil {
		t.Errorf("Expected Logging to be set, got nil")
	}
}
