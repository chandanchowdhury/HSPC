package dbhandler

import (
	"database/sql"
	"fmt"
	"log"
)

//TODO: Read from a config file or environment
const (
	DB_USER     = "hspc"
	DB_PASSWORD = "HSPC-Password"
	DB_NAME     = "postgres"
)

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Print("Error connecting DB")
		log.Fatal(err)
	}
	return db
}
