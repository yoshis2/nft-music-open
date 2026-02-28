// Package util は、共通のユーティリティ関数を提供します。
package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// JapaneseNowTimeのテスト
func TestJapaneseNowTime(t *testing.T) {
	nowJST := JapaneseNowTime()

	// 日本時間のタイムゾーンを取得
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	// 現在の日本時間を計算
	expectedTime := time.Now().In(jst)

	// 取得した時間が日本時間の範囲内であることを確認
	assert.WithinDuration(t, nowJST, expectedTime, time.Second, "JapaneseNowTime should return the current time in JST")
}
