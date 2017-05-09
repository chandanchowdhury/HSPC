package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

//TODO: Read from a config file or environment
const (
	DB_USER     = "hspc"
	DB_PASSWORD = "HSPC-Password"
	DB_NAME     = "postgres"

	foreignKeyViolationErrorCode = pq.ErrorCode("23503")
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

func isForeignKeyError(err error) bool {

	if pgErr, isPGErr := err.(*pq.Error); isPGErr {
		log.Printf("PostgreSQL Error Code: %s", pgErr.Code)
		if pgErr.Code == foreignKeyViolationErrorCode {
			// handle foreign_key_violation errors here
			log.Print("Foreign Key Violation")
			return true
		}
	}

	return false
}
