package dbhandler

import (
	"database/sql"
	"log"

	"github.com/chandanchowdhury/HSPC/models"
)

/**
Credential
 **/
func CredentialCreate(credential models.Credential) int64 {
	log.Print("# Creating credential")

	db := getDBConn()

	stmt, err := db.Prepare("INSERT INTO Credential(emailaddress, password_hash)" +
		" VALUES($1, $2) returning credential_id")
	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var lastInsertId int64
	err = stmt.QueryRow(credential.Emailaddress, credential.Password).Scan(&lastInsertId)

	if err != nil {
		if isDuplicateKeyError(err) {
			return -1
		}

		log.Panic(err)
	}

	log.Printf("credential_id = %d", lastInsertId)

	return lastInsertId
}

func CredentialRead(emailaddress string) *models.Credential {
	log.Print("# Reading Credential")
	log.Printf("emailaddress = %s", emailaddress)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT credential_id, emailaddress, password_hash, credential_active " +
		"FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	credential := new(models.Credential)
	err = stmt.QueryRow(emailaddress).Scan(
		&credential.CredentialID,
		&credential.Emailaddress,
		&credential.Password,
		&credential.CredentialActive)

	// if no records found, return an empty struct instead of failing
	if err == sql.ErrNoRows {
		return &models.Credential{}
	}

	if err != nil {
		log.Print("Error reading Credential data")
		log.Panic(err)
	}

	return credential
}

func CredentialUpdate(emailaddress string, password string, credential_active bool) int64 {
	db := getDBConn()

	log.Print("# Updating Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("UPDATE Credential SET emailaddress = $1, " +
		"password_hash = $2, credential_active = $3 WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(emailaddress, password, credential_active)

	if err != nil {
		if isForeignKeyError(err) {
			return -1
		}

		if isDuplicateKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func CredentialDelete(emailaddress string) int64 {
	db := getDBConn()

	log.Print("# Deleting Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("DELETE FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(emailaddress)

	if err != nil {
		log.Print("Error deleting from Credential")

		if isForeignKeyError(err) {
			return -2
		}

		log.Print(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}
