package dbhandler

import "log"

/**
  SchoolAddAdvisor adds an Advisor to a School
*/
func SchoolAddAdvisor(school_id int64, advisor_id int64) bool {
	log.Print("# Add School Advisor")
	log.Printf("School ID = %d, Advisor ID = %d", school_id, advisor_id)

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO School_Advisor(school_id, advisor_id) " +
		"VALUES($1, $2)")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := stmt.Exec(school_id, advisor_id)

	if err != nil {
		if isForeignKeyError(err) {
			return false
		}

		if isDuplicateKeyError(err) {
			return false
		}

		log.Panic(err)
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows != 1 {
		log.Panic("Unexpected number of inserts")
	}

	return true
}

/**
SchoolDeleteAdvisor deletes an Advisor from a School
*/
func SchoolDeleteAdvisor(school_id int64, advisor_id int64) bool {
	log.Print("# Remove School Advisor")
	log.Printf("School ID = %d, Asvisor ID = %d", school_id, advisor_id)

	db := getDBConn()

	delete_stmt, err := db.Prepare("DELETE FROM School_Advisor WHERE school_id = $1 AND advisor_id = $2")
	defer delete_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := delete_stmt.Exec(school_id, advisor_id)

	affectedRows, err := result.RowsAffected()
	if affectedRows != 1 {
		if affectedRows == 0 {
			return false
		}
		log.Panic("Unexpected number of deletion")
	}

	if err != nil {
		if isForeignKeyError(err) {
			return false
		}

		if isDuplicateKeyError(err) {
			return false
		}

		log.Panic(err)
	}

	return true
}

/**
Given a school_id, get the advisor for the School.
*/
func SchoolReadAdvisor(school_id int64) int64 {
	log.Print("# Reading School Advisor")
	log.Printf("School ID = %d", school_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT advisor_id " +
		"FROM School_Advisor WHERE school_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var advisor_id int64
	err = stmt.QueryRow(school_id).Scan(&advisor_id)

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	return advisor_id
}

/**
Given an AdvisorID, find all Schools
*/
func AdvisorGetAllSchools(advisor_id int64) []int64 {
	log.Printf("Read all Schools for AdvisorID = %d", advisor_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT school_id " +
		"FROM School_Advisor WHERE advisor_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	crsr, err := stmt.Query(advisor_id)

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	schools := make([]int64, 0)
	for crsr.Next() {
		var school_id int64
		crsr.Scan(&school_id)
		schools = append(schools, school_id)
	}

	return schools
}
