// Package util は、共通のユーティリティ関数を提供します。
package util

import "database/sql"

func EmptyString(value sql.NullString) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}
