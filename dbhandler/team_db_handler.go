package dbhandler

import (
	"database/sql"
	"github.com/chandanchowdhury/HSPC/models"
	"log"
)

const (
	MAX_STUDENT_PER_TEAM = 4
)

/*
Team
*/
func TeamCreate(team models.Team) int64 {
	log.Print("# Creating Team")

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO team(team_name, team_division, school_id) " +
		"VALUES($1, $2, $3) returning team_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var team_id int64
	err = stmt.QueryRow(team.TeamName, team.TeamDivision, team.SchoolID).Scan(&team_id)

	//log.Print(err)
	//log.Print(pq.ErrorClass("foreign_key_violation"))
	//TODO: Check FK constraint violation
	//if err.Error() ==  pq.ErrorClass("foreign_key_violation").Name() {
	//log.Print("Foreign Key Violation")
	//return -1;
	//}

	if err != nil {
		log.Print(err)
		return -1
	}

	//if TeamMember is not empty
	if len(team.TeamMembers) > 0 {
		//Add all the students to the Team
		for _, s_id := range team.TeamMembers {
			success := TeamAddMember(team.TeamID, s_id)

			if !success {
				log.Panic("Error adding student to team")
			}
		}
	}

	return team_id
}

/*
   Add a Student to a team
*/
func TeamAddMember(team_id int64, student_id int64) bool {

	log.Print("# Add member to a Team")

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO TeamStudent(team_id, student_id) " +
		"VALUES($1, $2)")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
		return false
	}

	result, err := stmt.Exec(team_id, student_id)

	if err != nil {
		log.Print(err)
		return false
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows != 1 {
		log.Print("Unexpected number of inserts")
		return false
	}

	//TODO: Check FK constraint violation
	//if err.Error() ==  pq.ErrorClass("foreign_key_violation").Name() {
	//log.Print("Foreign Key Violation")
	//return -1;
	//}

	return true
}

func TeamRead(team_id int64) models.Team {
	log.Print("# Reading Address")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT team_id, team_name, team_division, school_id " +
		"FROM team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	var team = models.Team{}
	err = stmt.QueryRow(team_id).Scan(&team.TeamID, &team.TeamName, &team.TeamDivision, &team.SchoolID)

	// if no records found, return an empty struct
	if err == sql.ErrNoRows {
		return models.Team{}
	}

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	//fill the members
	team.TeamMembers = TeamReadMembers(team_id)

	return team
}

func TeamReadMembers(team_id int64) []int64 {
	var student_ids []int64

	log.Print("# Reading Address")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT team_id, student_id " +
		"FROM TeamStudent WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	crsr, err := stmt.Query(team_id)

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	// if no records found, return an empty struct
	//if err == sql.ErrNoRows {
	//    return student_ids
	//}

	var s_id int64
	for crsr.Next() {
		//NOTE: Possibility of bug from UI handling
		if len(student_ids) == 0 {
			student_ids = make([]int64, 1, MAX_STUDENT_PER_TEAM)
		}
		crsr.Scan(&s_id)
		student_ids = append(student_ids, s_id)
	}

	return student_ids
}

func TeamUpdate(team models.Team) int64 {
	db := getDBConn()

	log.Print("# Updating Team")
	log.Printf("Team ID = %d", team.TeamID)

	stmt, err := db.Prepare("UPDATE team SET team_name = $1, " +
		"team_division = $2, school_id = $3 WHERE team_id = $4")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	//transaction begin
	tx, err := db.Begin()

	result, err := stmt.Exec(team.TeamName, team.TeamDivision, team.SchoolID, team.TeamID)

	if err != nil {
		tx.Rollback()
		log.Print("Error updating team")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	//we should update only one record
	if affectedCount > 1 {
		// rollback the update
		tx.Rollback()
		log.Panic("Unexpected number of updates: %d", affectedCount)
	}

	//commit the update
	tx.Commit()

	return affectedCount
}

func TeamMemberUpdate(team_id int64, student_ids []int64) bool {
	db := getDBConn()

	log.Print("# Update Team Membrs")
	log.Printf("Team ID = %d", team_id)

	//first delete all entries
	delete_stmt, err := db.Prepare("DELETE FROM TeamStudent WHERE team_id = $1")
	defer delete_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	insert_stmt, err := db.Prepare("DELETE FROM TeamStudent WHERE team_id = $1")
	defer insert_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	//transaction begin
	tx, err := db.Begin()

	_, err = delete_stmt.Exec(team_id)

	if err != nil {
		tx.Rollback()
		log.Print("Delete Failed")
		log.Panic(err)
	}

	for _, s_id := range student_ids {
		success := TeamAddMember(team_id, s_id)

		if !success {
			tx.Rollback()
			log.Panic("Error adding student to team for TeamMemberUdate")
		}
	}

	//commit the deletion
	tx.Commit()

	//then re-add all members
	return true
}

func TeamDelete(team_id int64) int64 {
	db := getDBConn()

	log.Print("# Deleting Team")
	log.Printf("Team ID = %d", team_id)

	stmt, err := db.Prepare("DELETE FROM team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	//transaction begin
	tx, err := db.Begin()

	result, err := stmt.Exec(team_id)

	if err != nil {
		tx.Rollback()
		log.Print("Delete Failed")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	//we should update only one record
	if affectedCount != 1 {
		tx.Rollback()
		log.Printf("Unexpected number of updates: %d", affectedCount)
	}

	//commit the deletion
	tx.Commit()

	return affectedCount
}
