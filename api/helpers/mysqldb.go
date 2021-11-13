package helpers

import (
	"database/sql"
)

func DbConn() (db *sql.DB) {
	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "root"
	// dbName := "todo"
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	db, err := sql.Open("mysql", "misuser:misuser123@tcp(10.2.10.154:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	return db
}
