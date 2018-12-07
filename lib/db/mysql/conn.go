package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Open(host string) error {
	db, err := sql.Open("mysql", host)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	con = db

	log.Println("mysql connected")
	return nil
}

func Close() error {
	if err := con.Close(); err != nil {
		return err
	}
	con = nil
	return nil
}

func DB() *sql.DB {
	return con
}
