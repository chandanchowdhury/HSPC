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

	foreignKeyViolationErrorCode   = pq.ErrorCode("23503")
	duplicateKeyViolationErrorCode = pq.ErrorCode("23505")
)

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Print("Error opening DB Connection")
		log.Fatal(err)
	}
	//check if the DB server is alive
	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		//if DB server is not alive, we cannot do anything, hence die die die...
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

func isDuplicateKeyError(err error) bool {

	if pgErr, isPGErr := err.(*pq.Error); isPGErr {
		log.Printf("PostgreSQL Error Code: %s", pgErr.Code)
		if pgErr.Code == duplicateKeyViolationErrorCode {
			// handle foreign_key_violation errors here
			log.Print("Duplicate Key Violation")
			return true
		}
	}

	return false
}
