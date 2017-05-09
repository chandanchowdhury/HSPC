package dbhandler

import (
	"database/sql"
	"github.com/chandanchowdhury/HSPC/models"
	"log"
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
		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	log.Printf("New Advisor ID: %d", advisor_id)
	return advisor_id
}

func AdvisorRead(advisor_id int64) models.Advisor {
	var advisor = models.Advisor{}

	log.Print("# Reading Advisor")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT advisor_id, advisor_name, credential_id " +
		"FROM advisor WHERE advisor_id = $1")
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
	log.Print("# Updating Advisor")
	log.Printf("Advisor ID = %d", advisor.AdvisorID)

	db := getDBConn()

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

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func AdvisorDelete(advisor_id int64) int64 {
	log.Print("# Deleting Advisor")
	log.Printf("Advisor ID = %d", advisor_id)

	db := getDBConn()

	stmt, err := db.Prepare("DELETE FROM advisor WHERE advisor_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(advisor_id)

	if err != nil {
		log.Print(err)

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func AdvisorReadAll() []*models.Advisor {
	log.Print("# Reading All Advisors")

	db := getDBConn()

	stmt, err := db.Prepare("SELECT advisor_id, advisor_name, credential_id " +
		"FROM advisor")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	crsr, err := stmt.Query()

	if err != nil {
		log.Print("Error getting advisor data")
		log.Panic(err)
	}

	advisors := make([]*models.Advisor, 0)
	for crsr.Next() {
		var advisor = models.Advisor{}
		crsr.Scan(&advisor.AdvisorID, &advisor.AdvisorName, &advisor.CredentialID)
		advisors = append(advisors, &advisor)
	}

	return advisors
}
