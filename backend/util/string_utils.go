// Package util は、共通のユーティリティ関数を提供します。
package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NormalizeAndFold は文字列を小文字に変換し、検索に適した形に正規化します。
// 主に大文字小文字を区別しない比較や検索に使用されます。
func NormalizeAndFold(s string) string {
	// Unicode正規化（NFKC）を行い、結合文字などを標準形にする
	// このプロジェクトでは特定のUnicode正規化は不要と判断し、strings.ToLowerとFoldを使用
	// 必要であればgolang.org/x/text/secure/precisやgolang.org/x/text/transformパッケージを検討

	// 大文字小文字を区別しない比較のために小文字に変換
	// strings.ToLowerはASCII文字だけでなくUnicode文字も適切に小文字に変換する
	// strings.ToLowerを使用する方が、strings.EqualFoldの内部処理に近く、一貫性がある
	return strings.ToLower(s)
}

// ContainsFold は s が substr を大文字小文字を区別せずに含むかどうかを報告します。
func ContainsFold(s, substr string) bool {
	// 両方の文字列を正規化（小文字に変換）してから部分文字列を検索する
	return strings.Contains(NormalizeAndFold(s), NormalizeAndFold(substr))
}

// Capitalize は文字列の最初の文字を大文字にし、残りの文字を小文字にします。
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	return cases.Title(language.Und).String(strings.ToLower(s))
}
