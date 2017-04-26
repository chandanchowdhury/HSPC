package dbhandler

import (
    "log"
    "github.com/chandanchowdhury/HSPC/models"
    "database/sql"
)

/*
School
*/
func SchoolCreate(school models.School) int64 {
    log.Print("# Creating School")

    db := getDBConn()
    stmt, err := db.Prepare("INSERT INTO school(school_name, address_id) " +
        "VALUES($1, $2) returning school_id")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    var school_id int64
    err = stmt.QueryRow(school.SchoolName, school.AddressID).Scan(&school_id)

    if err != nil {
        log.Print(err)
        return 0
    }

    return school_id
}

func SchoolRead(school_id int64) models.School {
    log.Print("# Reading Address")

    db := getDBConn()
    stmt, err := db.Prepare("SELECT school_id, school_name, address_id " +
        "FROM school WHERE school_id = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    var school = models.School{}
    err = stmt.QueryRow(school_id).Scan(&school.SchoolID, &school.SchoolName, &school.AddressID)

    // if no records found, return an empty struct
    if err == sql.ErrNoRows {
        return models.School{}
    }

    if err != nil {
        log.Print("Error getting school data")
        log.Panic(err)
    }

    return school
}

func SchoolUpdate(school models.School) int64 {
    db := getDBConn()

    log.Print("# Updating School")
    log.Printf("School ID = %d", school.SchoolID)

    stmt, err := db.Prepare("UPDATE school SET school_name = $1, " +
        "address_id = $2 WHERE school_id = $3")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    result, err := stmt.Exec(school.SchoolName, school.AddressID, school.SchoolID)

    if err != nil {
        log.Print("Error updating school")
        log.Panic(err)
    }

    affectedCount, err := result.RowsAffected()

    if affectedCount != 1 {
        log.Printf("Unexpected number of updates: %d", affectedCount)
    }

    return affectedCount
}

func SchoolDelete(school_id int64) int64 {
    db := getDBConn()

    log.Print("# Deleting School")
    log.Printf("School ID = %d", school_id)

    stmt, err := db.Prepare("DELETE FROM school WHERE school_id = $1")
    defer stmt.Close()

    if err != nil {
        log.Print("Error creating prepared statement")
        log.Print(err)
    }

    result, err := stmt.Exec(school_id)

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
