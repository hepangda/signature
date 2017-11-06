package main

import (
	"database/sql"
)

type signature struct {
	ID       string `json:"id"`
	Username string `json:"name"`
}

type action interface {
	Do(conn *sql.DB) error
}

type actRecord struct {
	Dat signature
}
