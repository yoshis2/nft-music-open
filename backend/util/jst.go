// Package util は、共通のユーティリティ関数を提供します。
package util

import "time"

// JapaneseNowTime は日本の現在時間を取得
func JapaneseNowTime() time.Time {
	// Time Zone を日本時間に設定
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().UTC().In(jst)
	return nowJST
}
