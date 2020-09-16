package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var DbConnection *sqlx.DB

func DbConnect () error {
	db, err := sqlx.Connect("mysql", os.Getenv("DATABASE_URL"))
	DbConnection = db
	return err
}