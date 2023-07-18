package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (a *App) ConnectDB() {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/my_go_db?multiStatements=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	a.DB = db
}
