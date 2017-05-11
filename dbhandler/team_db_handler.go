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
	log.Print("Creating Team")

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

	if err != nil {
		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
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

	log.Printf("New team_id = %d", team_id)

	return team_id
}

func TeamRead(team_id int64) models.Team {
	log.Printf("# Reading TeamID: %d", team_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT team_id, team_name, team_division, school_id " +
		"FROM team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
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

func TeamUpdate(team models.Team) int64 {
	log.Printf("Updating TeamID = %d", team.TeamID)

	db := getDBConn()

	stmt, err := db.Prepare("UPDATE team SET team_name = $1, " +
		"team_division = $2, school_id = $3 WHERE team_id = $4")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	//transaction begin
	//tx, err := db.Begin()

	result, err := stmt.Exec(team.TeamName, team.TeamDivision, team.SchoolID, team.TeamID)

	if err != nil {
		//tx.Rollback()
		log.Print("Error updating team")

		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()
	//we should update only one record
	if affectedCount != 1 {
		// rollback the update
		//tx.Rollback()

		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	//commit the update
	//tx.Commit()

	return affectedCount
}

func TeamDelete(team_id int64) int64 {
	log.Printf("Deleting Team ID = %d", team_id)

	db := getDBConn()

	stmt, err := db.Prepare("DELETE FROM team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	//transaction begin
	//tx, err := db.Begin()

	result, err := stmt.Exec(team_id)

	if err != nil {
		//tx.Rollback()
		log.Print("Delete Failed")

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	//we should update only one record
	if affectedCount != 1 {
		//tx.Rollback()

		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	//commit the deletion
	//tx.Commit()

	return affectedCount
}
