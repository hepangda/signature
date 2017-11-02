package main

//TODO: Add code this file
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func dbInit() error {
	dbConn, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/signature")
	return err
}

func dbRecord() error {
	//TODO
	return nil
}
