package config

import (
	"database/sql"
	"os"
	"log"
	
	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "exampleDB",
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db, nil
}