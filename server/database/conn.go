package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	c := mysql.Config{
		DBName:    "connect_todo",
		User:      "n000r111",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
