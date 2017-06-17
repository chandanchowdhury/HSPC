package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"os"
)

var (
	POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
	POSTGRES_DB       = os.Getenv("POSTGRES_DB")
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")

	foreignKeyViolationErrorCode   = pq.ErrorCode("23503")
	duplicateKeyViolationErrorCode = pq.ErrorCode("23505")
)

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD)

	log.Printf("PostgresSQL Host = %s", POSTGRES_HOST)
	log.Print(dbinfo)

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

	// list the available databases
	//dbs, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
	//
	//var dbname string
	//for dbs.Next() {
	//	dbs.Scan(&dbname)
	//	log.Print(&dbname)
	//}

	// list the available tables
	//tables, err := db.Query("SELECT table_schema,table_name FROM information_schema.tables ORDER BY table_schema,table_name")
	//var tschema, tname string
	//for tables.Next() {
	//	tables.Scan(&tschema, &tname)
	//	log.Print(tschema, tname)
	//}

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
