package main

import (
	"fmt"
	"log"

	"github.com/h3th-IV/mySQL/db"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("mySQL")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.Init()
	defer db.Close()
	db.Insert()
}
