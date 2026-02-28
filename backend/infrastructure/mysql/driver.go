// Package mysql は、MySQLデータベースへの接続と操作を提供します。
package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql dbをmysqlにする場合はこれを使う
type Mysql struct{}

// NewMysql New Mysql
func NewMysql() *Mysql {
	return &Mysql{}
}

// Open はGormConnect MySQL wrapper に接続
func (mysqls *Mysql) Open() *gorm.DB {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	option := "charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, password, host, name, option)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
