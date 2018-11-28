package psql

import "database/sql"

var conn *sql.DB

//open connection for postgresql
func Open(host string) error {
	db, err := sql.Open("postgres", host)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	conn = db

	return nil
}

//close connection for postgresql
func Close() error {
	if err := conn.Close(); err != nil {
		return err
	}
	conn = nil
	return nil
}

func DB() *sql.DB {
	return conn
}
