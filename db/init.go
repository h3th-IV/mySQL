package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func Init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MY_USER"), os.Getenv("MY_PASSWORD"), os.Getenv("MY_HOST"), os.Getenv("MY_PORT"), os.Getenv("MY_DBNAME"))

	var err error
	dB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = dB.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connection to Database was succeefull")
}

func Close() {
	dB.Close()
}
