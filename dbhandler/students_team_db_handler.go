package dbhandler

import "log"

/*
   Add a Student to a team
*/
func TeamAddMember(team_id int64, student_id int64) bool {
	log.Print("# Add Team member")
	log.Printf("Team ID = %d, Student ID = %d", team_id, student_id)

	db := getDBConn()

	//TODO: Make sure the Team and Student belong to the same School

	stmt, err := db.Prepare("INSERT INTO Student_Team(team_id, student_id) " +
		"VALUES($1, $2)")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := stmt.Exec(team_id, student_id)

	if err != nil {
		if isForeignKeyError(err) {
			return false
		}

		if isDuplicateKeyError(err) {
			//if the entry already exists
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

func TeamDeleteMember(team_id int64, student_id int64) bool {
	log.Print("# Remove Team Member")
	log.Printf("Team ID = %d, Student ID = %d", team_id, student_id)

	db := getDBConn()

	delete_stmt, err := db.Prepare("DELETE FROM Student_Team WHERE team_id = $1 AND student_id = $2")
	defer delete_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := delete_stmt.Exec(team_id, student_id)

	if err != nil {
		log.Print("Delete Failed")

		if isForeignKeyError(err) {
			return false
		}

		log.Panic(err)
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows != 1 {
		log.Panic("Unexpected number of deletion")
	}

	return true
}

func TeamReadMembers(team_id int64) []int64 {
	log.Print("# Reading Team members")
	log.Printf("Team ID = %d", team_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT student_id " +
		"FROM Student_Team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	crsr, err := stmt.Query(team_id)

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	student_ids := make([]int64, 0, MAX_STUDENT_PER_TEAM)
	var s_id int64
	for crsr.Next() {
		crsr.Scan(&s_id)
		student_ids = append(student_ids, s_id)
	}

	return student_ids
}
