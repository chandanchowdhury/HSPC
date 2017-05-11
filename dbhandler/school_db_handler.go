package dbhandler

import (
	"database/sql"
	"log"

	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
)

/*
School
*/
func SchoolCreate(school models.School) int64 {
	log.Print("Creating School")

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO school(school_name, address_id) " +
		"VALUES($1, $2) returning school_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var school_id int64
	err = stmt.QueryRow(school.SchoolName, school.AddressID).Scan(&school_id)

	if err != nil {
		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	log.Printf("New school_id = %d", school_id)
	return school_id
}

func SchoolRead(school_id int64) models.School {
	log.Printf("Reading school_id = %d", school_id)

	db := getDBConn()
	stmt, err := db.Prepare("SELECT school_id, school_name, address_id " +
		"FROM school WHERE school_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
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
	log.Printf("Updating School ID = %d", school.SchoolID)

	db := getDBConn()

	stmt, err := db.Prepare("UPDATE school SET school_name = $1, " +
		"address_id = $2 WHERE school_id = $3")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	//tx, err := db.Begin()
	result, err := stmt.Exec(school.SchoolName, school.AddressID, school.SchoolID)

	if err != nil {
		//tx.Rollback()
		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount > 1 {
		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

		// rollback the update
		//tx.Rollback()
		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	//commit the update
	//tx.Commit()

	return affectedCount
}

func SchoolDelete(school_id int64) int64 {
	log.Printf("Deleting School ID = %d", school_id)

	db := getDBConn()

	stmt, err := db.Prepare("DELETE FROM school WHERE school_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := stmt.Exec(school_id)

	if err != nil {
		if isForeignKeyError(err) {
			return -2
		}

		log.Print("Delete Failed")
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

func SchoolList(params school.GetSchoolParams) []*models.School {
	log.Print("# Reading School List")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT school_id, school_name, address_id FROM school")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	crsr, err := stmt.Query()

	// if no records found, return an empty struct
	if err == sql.ErrNoRows {
		return []*models.School{}
	}

	if err != nil {
		log.Print("Error getting school data")
		log.Panic(err)
	}

	// create an array of size zero
	school_list := make([]*models.School, 0)
	//till there are records
	for crsr.Next() {
		// create a new object to hold the data
		school := new(models.School)
		// fetch the next record into the new object
		crsr.Scan(&school.SchoolID, &school.SchoolName, &school.AddressID)
		//append to the list
		school_list = append(school_list, school)
	}

	return school_list
}
