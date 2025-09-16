package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MYSQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MYSQLStorage {
	db , err := sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL!")

	return &MYSQLStorage{db: db}
}