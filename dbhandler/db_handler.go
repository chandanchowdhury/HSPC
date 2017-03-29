package dbhandler

import (
	"fmt"
	"database/sql"
	// the driver is used internally, the underscore makes sure the "unused"
	// error is suppressed.
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "hspc"
	DB_PASSWORD = "HSPC-Password"
	DB_NAME     = "postgres"
)

func checkErr(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

/**
Insert an email address into the table
 **/
func EmailAddressCreate(emailaddress string) int {
	db := getDBConn()

	log.Printf("# Inserting emailaddress")

	var lastInsertId int
	err := db.QueryRow("INSERT INTO emailaddress(emailaddress) VALUES($1) returning emailaddress_id;", emailaddress).Scan(&lastInsertId)
	checkErr(err)
	log.Printf("emailaddress_id = %d", lastInsertId)

	return lastInsertId
}

/**
Read an email address from the table
 **/
func EmailAddressRead(emailaddress_id int) string {
	var emailaddress string
	db := getDBConn()

	log.Printf("# Reading emailaddress")

	rows := db.QueryRow("SELECT emailaddress FROM emailaddress WHERE emailaddress_id = $1", emailaddress_id)

	err := rows.Scan(&emailaddress)
	checkErr(err)

	log.Printf("emailaddress_id = %d, emailaddress = %s", emailaddress_id, emailaddress)

	return emailaddress
}


/**
Update an email address in the table
 **/
func EmailAddressUpdate(emailaddress_id int, emailaddress string) int {
	db := getDBConn()

	log.Printf("# Updating emailaddress")
	log.Printf("emailaddress_id = %d, emailaddress = %s", emailaddress_id, emailaddress)

	stmt, err := db.Prepare("UPDATE emailaddress SET emailaddress = $1 WHERE emailaddress_id = $2")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress, emailaddress_id)

	checkErr(err)

	updateCount, err := result.RowsAffected()

	if ( updateCount == 1) {
		return emailaddress_id
	}

	return -1
}

/**
Delete an email address in the table
 **/
func EmailAddressDelete(emailaddress_id int) int {
	db := getDBConn()

	log.Printf("# Deleting emailaddress")
	log.Printf("emailaddress_id = %d", emailaddress_id)

	stmt, err := db.Prepare("DELETE FROM emailaddress WHERE emailaddress_id = $1")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress_id)

	checkErr(err)

	deleteCount, err := result.RowsAffected()

	if ( deleteCount == 1) {
		return emailaddress_id
	}

	return -1
}
