package dbhandler

import (
    "log"
    "github.com/chandanchowdhury/HSPC/models"
    "database/sql"
)

/*
Advisor
*/
func AdvisorCreate(advisor models.Advisor) int64 {
    log.Print("# Creating Advisor")

    db := getDBConn()
    stmt, err := db.Prepare("INSERT INTO advisor(advisor_name, credential_id) " +
        "VALUES($1, $2) returning advisor_id")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    var advisor_id int64
    err = stmt.QueryRow(advisor.AdvisorName, advisor.CredentialID).Scan(&advisor_id)

    if err != nil {
        log.Print(err)
        return 0
    }

    return advisor_id
}

func AdvisorRead(advisor_id int64) models.Advisor {
    var advisor = models.Advisor{}

    log.Print("# Reading Advisor")

    db := getDBConn()
    stmt, err := db.Prepare("SELECT advisor_id, advisor_name, credential_id " +
        "FROM advisor WHERE advior_id = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    err = stmt.QueryRow(advisor_id).Scan(&advisor.AdvisorID, &advisor.AdvisorName, &advisor.CredentialID)

    // if no records found, return an empty struct
    if err == sql.ErrNoRows {
        return models.Advisor{}
    }

    if err != nil {
        log.Print("Error getting advisor data")
        log.Panic(err)
    }

    return advisor
}

func AdvisorUpdate(advisor models.Advisor) int64 {
    db := getDBConn()

    log.Print("# Updating Advisor")
    log.Printf("Advisor ID = %d", advisor.AdvisorID)

    stmt, err := db.Prepare("UPDATE advisor SET advisor_name = $1, credential_id = $2" +
        "WHERE advisor_id = $3")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    result, err := stmt.Exec(advisor.AdvisorName, advisor.CredentialID, advisor.AdvisorID)

    if err != nil {
        log.Print("Error updating advisor")
        log.Panic(err)
    }

    affectedCount, err := result.RowsAffected()

    if affectedCount != 1 {
        log.Printf("Unexpected number of updates: %d", affectedCount)
    }

    return affectedCount
}

func AdvisorDelete(advisor_id int64) int64 {
    db := getDBConn()

    log.Print("# Deleting Advisor")
    log.Printf("Advisor ID = %d", advisor_id)

    stmt, err := db.Prepare("DELETE FROM advisor WHERE advisor_id = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    result, err := stmt.Exec(advisor_id)

    if err != nil {
        log.Print("Delete Failed")
        log.Panic(err)
    }

    affectedCount, err := result.RowsAffected()

    if affectedCount != 1 {
        log.Printf("Unexpected number of updates: %d", affectedCount)
    }

    return affectedCount
}
