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
func CredentialCreate(emailaddress string, password_hash string) uint32 {
	db := getDBConn()

	log.Printf("# Creating credential")

	var lastInsertId uint32

	err := db.QueryRow("INSERT INTO Credential(emailaddress, password_hash) VALUES($1, $2) returning credential_id;", emailaddress, password_hash).Scan(&lastInsertId)
	checkErr(err)
	log.Printf("credential_id = %d", lastInsertId)

	return lastInsertId
}

/**
Read an credential from the table
 **/
func CredentialRead(emailaddress string) credential_struct {
	var credential = credential_struct{}
	db := getDBConn()

	log.Printf("# Reading Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("SELECT credential_id, emailaddress, password_hash FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	rows, err := stmt.Query(emailaddress)
	defer rows.Close()

	checkErr(err)

	for rows.Next() {
		err := rows.Scan(&credential.credential_id, &credential.emailaddress, &credential.password_hash)
		checkErr(err)
		return credential
	}

	return credential
}

/**
Update an email address in the table
 **/
func CredentialUpdate(emailaddress string, password string) int32 {
	db := getDBConn()

	log.Printf("# Updating Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("UPDATE Credential SET emailaddress = $1, password_hash = $2 WHERE emailaddress = $1")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress, password)

	checkErr(err)

	updateCount, err := result.RowsAffected()

	if ( updateCount == 1) {
		return 0
	}

	return -1
}

/**
Delete an email address from the table
 **/
func CredentialDelete(emailaddress string) uint32 {
	db := getDBConn()

	log.Printf("# Deleting Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("DELETE FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress)

	checkErr(err)

	deleteCount, err := result.RowsAffected()

	if ( deleteCount == 1) {
		return 0
	}

	return 0
}

func AddressCreate(address address_struct) uint32 {

    //TODO: complete the logic

    return 0
}

func AddressRead(address_id uint32) address_struct {
    var address = address_struct{}

    //TODO: complete the logic

    return address
}

func AddressUpdate(address address_struct) address_struct {
    var address = address_struct{}

    //TODO: complete the logic

    return address
}

func AddressDelete(address_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}