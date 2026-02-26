package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectDatabase() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3307)/inventory_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Koneksi Database Error: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	log.Println("Database Connected Succesfully")
}