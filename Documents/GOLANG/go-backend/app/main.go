package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"example.com/app/routes"
	"github.com/go-sql-driver/mysql"
)

func main() {
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

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }

	// fmt.Println("Connected!")

	router := routes.SetupRouter(db)
	router.Run("localhost:8080")
}