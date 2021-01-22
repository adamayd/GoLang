package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxIdleConns(4)
	DbConn.SetMaxOpenConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
