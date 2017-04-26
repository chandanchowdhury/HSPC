package dbhandler

import (
    "log"
    "github.com/chandanchowdhury/HSPC/models"
    "database/sql"
    "github.com/go-openapi/strfmt"
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
        log.Print(err)
    }

    var lastInsertId int64
    err = stmt.QueryRow(credential.Emailaddress, credential.Password).Scan(&lastInsertId)

    if err != nil {
        log.Print(err)
    }

    log.Printf("credential_id = %d", lastInsertId)

    return lastInsertId
}

func CredentialRead(emailaddress string) models.Credential {
    log.Print("# Reading Credential")
    log.Printf("emailaddress = %s", emailaddress)

    db := getDBConn()
    stmt, err := db.Prepare("SELECT credential_id, emailaddress, password_hash " +
        "FROM Credential WHERE emailaddress = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Panic(err)
    }

    var credential_id int64
    var email, pass string
    err = stmt.QueryRow(emailaddress).Scan(&credential_id, &email, &pass)

    if err == sql.ErrNoRows {
        return models.Credential{}
    }

    if err != nil {
        log.Print("Error reading Credential data")
        log.Panic(err)
    }

    log.Printf("DB Read - email: %s password: %s", email, pass)

    //setup the Credential model object
    var credential = models.Credential{}
    var e strfmt.Email = strfmt.Email(email)
    var p strfmt.Password = strfmt.Password(pass)
    credential.CredentialID = credential_id
    credential.Emailaddress = &e
    credential.Password = &p
    return credential
}

func CredentialUpdate(emailaddress string, password string) int64 {
    db := getDBConn()

    log.Print("# Updating Credential")
    log.Printf("emailaddress = %s", emailaddress)

    stmt, err := db.Prepare("UPDATE Credential SET emailaddress = $1, " +
        "password_hash = $2 WHERE emailaddress = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    result, err := stmt.Exec(emailaddress, password)

    if err != nil {
        log.Print("Error updating Credential")
        log.Print(err)
    }

    affectedCount, err := result.RowsAffected()

    if affectedCount != 1 {
        log.Printf("Unexpected number of updates: %d", affectedCount)
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
        log.Print(err)
    }

    affectedCount, err := result.RowsAffected()

    if affectedCount != 1 {
        log.Printf("Unexpected number of updates: %d", affectedCount)
    }

    return affectedCount
}
