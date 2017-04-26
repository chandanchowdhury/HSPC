package dbhandler

import (
	_ "database/sql"
	// the driver is used internally, the underscore makes sure the "unused"
	// error is suppressed.
	"github.com/chandanchowdhury/HSPC/models"
	_ "github.com/lib/pq"
	_ "gopkg.in/mgo.v2"
	_ "log"
)

/*
Team
*/
/*
func TeamCreate(team models.Team) int64 {
	log.Print("# Creating Team")

	db := getDBConn()

	team_stmt, err := db.Prepare("INSERT INTO team(team_name, team_division, school_id) " +
		"VALUES($1, $2, $3) returning team_id")
	defer team_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	//TODO: begin transaction
	//db.Begin()

	var team_id int64
	err = team_stmt.QueryRow(team.TeamName, team.TeamDivision, team.SchoolID).Scan(&team_id)

	if err != nil {
		//TODO: abort transaction
		log.Print("Error inserting into Team")
		log.Print(err)
		return 0
	}

	//insert the members into TeamStudent
	count := TeamStudentCreate(team)
	if count < count(team.TeamMembers) {
		log.Print("Not all students insrted into TeamStudent")
	}

	//TODO: commit

	return team_id
}

func TeamStudentCreate(team models.Team) int64 {
	log.Print("Creating TeamMembers")

	db := getDBConn()

	teamstudent_stmt, err := db.Prepare("INSERT INTO teamstudent(team_id, student_id) " +
		"VALUES($1, $2)")
	defer teamstudent_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	for _, student_id := range team.TeamMembers {
		teamstudent_stmt.Query(team.TeamID, student_id)

		if err != nil {
			log.Print("Error inserting into TeamStudent")
			log.Print(err)
			return 0
		}
	}

	// how many student in team now?
	stmt, err := db.Prepare("SELECT COUNT(student_id) FROM teamstudent WHERE team_id = $1")

	result := stmt.QueryRow(team.TeamID)

	var count int64
	result.Scan(&count)

	return count
}

func TeamRead(team_id int64) models.Team {
	var team = models.Team{}

	log.Print("# Reading Team")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT team_id, team_name, team_division, school_id " +
		"FROM team WHERE team_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	err = stmt.QueryRow(team_id).Scan(&team.TeamID, &team.TeamName, &team.TeamDivision, &team.SchoolID)

	// if no records found, return an empty struct
	if err == sql.ErrNoRows {
		return models.Team{}
	}

	if err != nil {
		log.Print("Error getting team data")
		log.Panic(err)
	}

	return team
}

func TeamStudentRead(team_id int64) []int64 {
	var members []int64

	db := getDBConn()

	// how many student in team now?
	stmt, err := db.Prepare("SELECT student_id FROM teamstudent WHERE team_id = $1")

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Query(team_id)

	//TODO: fill members array with values from query
	for result.Next() {

	}

	return members
}

func TeamUpdate(team models.Team) int64 {
	db := getDBConn()

	log.Print("# Updating Team")
	log.Printf("Team ID = %d", team.TeamID)

	stmt, err := db.Prepare("UPDATE team SET team_name = $1, team_division = $2, school_id = $3" +
		"WHERE team_id = $4")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(team.TeamName, team.TeamDivision, team.SchoolID, team.TeamID)

	if err != nil {
		log.Print("Error updating team")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Printf("Unexpected number of updates: %d", affectedCount)
	}

	//delete all previous team members and reinsert
	TeamStudentDelete(team.TeamID)
	insert_count := TeamStudentCreate(team)

	member_count := count(team.TeamMembers)

	if insert_count != member_count {
		log.Printf("Received %d but inserted %d", member_count, insert_count)
		//TODO: abort
	}

	//TODO: Commit

	return affectedCount
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

	//begin transaction
	result, err := stmt.Exec(team_id)

	if err != nil {
		log.Print("Delete Failed")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Printf("Unexpected number of updates: %d", affectedCount)
		//abort
	}

	count := TeamStudentDelete(team_id)
	log.Printf("Rows deleted from TeamStudent: %d", count)

	//commit

	return affectedCount
}

func TeamStudentDelete(team_id int64) int64 {
	log.Print("Deleting TeamMembers")

	db := getDBConn()

	teamstudent_stmt, err := db.Prepare("DELETE FROM teamstudent" +
		"WHERE team_id = $1")
	defer teamstudent_stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := teamstudent_stmt.Exec(team_id)
	if err != nil {
		log.Print("Error executing prepared statement")
		log.Print(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Print("Error getting affected rows")
		log.Print(err)
	}

	return count
}
*/

//Team Score

func teamscoreCreate(teamscore models.TeamScore) int64 {
	//TODO: complete the logic

	return 0
}

func teamscoreRead(teamscore_id int64) models.TeamScore {
	var teamscore = models.TeamScore{}

	//TODO: complete the logic

	return teamscore
}

func teamscoreUpdate(teamscore models.TeamScore) int64 {
	//TODO: complete the logic

	return 0
}

func teamscoreDelete(teamscore_id int64) int64 {
	//TODO: complete the logic

	return 0
}

/*
//Parking

func parkingCreate(parking models.Parking) int64 {
	//TODO: complete the logic

	return 0
}

func parkingRead(parking_id int64) models.Parking {
	var parking = models.Parking{}

	//TODO: complete the logic

	return parking
}

func parkingUpdate(parking models.Parking) int64 {
	//TODO: complete the logic

	return 0
}

func parkingDelete(parking_id int64) int64 {
	//TODO: complete the logic

	return 0
}
*/

/*
const (
	mongo_DB_USER     = "hspc"
	mongo_DB_PASSWORD = "HSPC-Password"
	mongo_DB_NAME     = "hspc"
)

func getMongoConn() *mgo.Database {
	mcred := mgo.Credential{mongo_DB_USER, mongo_DB_PASSWORD}
	db := mgo.Database{mongo_DB_NAME, mcred}

	return &db
}

//Problem

func problemCreate(problem models.Problem) int64 {
	//TODO: complete the logic

	db := getMongoConn()

	coll := db.C("problem")

	coll.Insert(problem)

	return 0
}

func problemRead(problem_id int64) models.Problem {
	var problem = models.Problem{}

	//TODO: complete the logic
	db := getMongoConn()
	coll := db.C("problem")

	query := mgo.Query{"problem_id": problem_id}

	result := coll.Find(query)

	if result.Count() > 1 {
		log.Printf("Expected 1 found %d", result.Count())
	}

	result.One(&problem)

	return problem
}

func problemUpdate(problem models.Problem) models.Problem {
	//TODO: complete the logic

	db := getMongoConn()
	coll := db.C("problem")

	query := mgo.Query{"problem_id": problem.ProblemID}
	err := coll.Update(query, problem)

	if err != nil {
		log.Print("Error updating Problem")
		log.Print(err)
	}

	return problem
}

func problemDelete(problem_id int64) int64 {
	//TODO: complete the logic
	db := getMongoConn()
	coll := db.C("problem")

	query := mgo.Query{"problem_id": problem_id}
	err := coll.RemoveId(query)

	if err != nil {
		log.Print("Error updating Problem")
		log.Print(err)
		return -1
	}

	return 0
}

//Solution

func solutionCreate(solution models.Solution) int64 {
	//TODO: complete the logic

	return 0
}

func solutionRead(solution_id int64) models.Solution {
	var solution = models.Solution{}

	//TODO: complete the logic

	return solution
}

func solutionUpdate(solution models.Solution) models.Solution {
	//TODO: complete the logic

	return solution
}

func solutionDelete(solution_id int64) int64 {
	//TODO: complete the logic

	return 0
}
*/
