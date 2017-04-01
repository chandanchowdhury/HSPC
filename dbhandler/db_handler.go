package dbhandler

import (
	"fmt"
	"database/sql"
	// the driver is used internally, the underscore makes sure the "unused"
	// error is suppressed.
	_ "github.com/lib/pq"
	"log"
)

//TODO: Read from a config file or environment
const (
	DB_USER     = "hspc"
	DB_PASSWORD = "HSPC-Password"
	DB_NAME     = "postgres"
)

func checkErr(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

/**
Credential
 **/
func CredentialCreate(emailaddress string, password_hash string) uint32 {
	db := getDBConn()

	log.Printf("# Creating credential")

	var lastInsertId uint32

	err := db.QueryRow("INSERT INTO Credential(emailaddress, password_hash) VALUES($1, $2) returning credential_id;", emailaddress, password_hash).Scan(&lastInsertId)
	if err != nil {
        return 0
    }

	log.Printf("credential_id = %d", lastInsertId)

	return lastInsertId
}

func CredentialRead(emailaddress string) credential_struct {
	var credential = credential_struct{}
	db := getDBConn()

	log.Printf("# Reading Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("SELECT credential_id, emailaddress, password_hash FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	rows, err := stmt.Query(emailaddress)
	defer rows.Close()

	checkErr(err)

	for rows.Next() {
		err := rows.Scan(&credential.credential_id, &credential.emailaddress, &credential.password_hash)
		checkErr(err)
		return credential
	}

	return credential
}

func CredentialUpdate(emailaddress string, password string) int32 {
	db := getDBConn()

	log.Printf("# Updating Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("UPDATE Credential SET emailaddress = $1, password_hash = $2 WHERE emailaddress = $1")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress, password)

	checkErr(err)

	updateCount, err := result.RowsAffected()

	if ( updateCount == 1) {
		return 0
	}

	return -1
}

func CredentialDelete(emailaddress string) uint32 {
	db := getDBConn()

	log.Printf("# Deleting Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("DELETE FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	result, err := stmt.Exec(emailaddress)

	checkErr(err)

	deleteCount, err := result.RowsAffected()

	if ( deleteCount == 1) {
		return 0
	}

	return 0
}

/*
Address
*/
func AddressCreate(address address_struct) uint32 {

    //TODO: complete the logic

    return 0
}

func AddressRead(address_id uint32) address_struct {
    var address = address_struct{}

    //TODO: complete the logic

    return address
}

func AddressUpdate(address address_struct) address_struct {
    //TODO: complete the logic

    return address
}

func AddressDelete(address_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
School
*/
func SchoolCreate(school school_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func SchoolRead(school_id uint32) school_struct {
    var school = school_struct{}
    //TODO: complete the logic

    return school
}

func SchoolUpdate(school school_struct) school_struct {
    //TODO: complete the logic

    return school
}

func SchoolDelete(school_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Advisor
*/
func AdvisorCreate(advisor advisor_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func AdvisorRead(advisor_id uint32) advisor_struct {
    var advisor = advisor_struct{}

    //TODO: complete the logic

    return advisor
}

func AdvisorUpdate(advisor advisor_struct) advisor_struct {
    //TODO: complete the logic

    return advisor
}

func AdvisorDelete(advisor_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Team
*/
func TeamCreate(team team_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func TeamRead(team_id uint32) team_struct {
    var team = team_struct{}

    //TODO: complete the logic

    return team
}

func TeamUpdate(team team_struct) team_struct {
    //TODO: complete the logic

    return team
}

func TeamDelete(team_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Student
*/
func studentCreate(student student_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func studentRead(student_id uint32) student_struct {
    var student = student_struct{}

    //TODO: complete the logic

    return student
}

func studentUpdate(student student_struct) student_struct {
    //TODO: complete the logic

    return student
}

func studentDelete(student_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Team Score
*/
func teamscoreCreate(teamscore team_score_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func teamscoreRead(teamscore_id uint32) team_score_struct {
    var teamscore = team_score_struct{}

    //TODO: complete the logic

    return teamscore
}

func teamscoreUpdate(teamscore team_score_struct) team_score_struct {
    //TODO: complete the logic

    return teamscore
}

func teamscoreDelete(teamscore_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Parking
*/
func parkingCreate(parking team_score_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func parkingRead(parking_id uint32) team_score_struct {
    var parking = team_score_struct{}

    //TODO: complete the logic

    return parking
}

func parkingUpdate(parking team_score_struct) team_score_struct {
    //TODO: complete the logic

    return parking
}

func parkingDelete(parking_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Problem
*/
func problemCreate(problem problem_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func problemRead(problem_id uint32) problem_struct {
    var problem = problem_struct{}

    //TODO: complete the logic

    return problem
}

func problemUpdate(problem problem_struct) problem_struct {
    //TODO: complete the logic

    return problem
}

func problemDelete(problem_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Solution
*/
func solutionCreate(solution solution_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func solutionRead(solution_id uint32) solution_struct {
    var solution = solution_struct{}

    //TODO: complete the logic

    return solution
}

func solutionUpdate(solution solution_struct) solution_struct {
    //TODO: complete the logic

    return solution
}

func solutionDelete(solution_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}

/*
Problem_Solution
*/
func problemsolutionCreate(problemsolution problem_solution_struct) uint32 {
    //TODO: complete the logic

    return 0
}

func problemsolutionRead(problemsolution_id uint32) problem_solution_struct {
    var problemsolution = problem_solution_struct{}

    //TODO: complete the logic

    return problemsolution
}

func problemsolutionUpdate(problemsolution problem_solution_struct) problem_solution_struct {
    //TODO: complete the logic

    return problemsolution
}

func problemsolutionDelete(problemsolution_id uint32) uint32 {
    //TODO: complete the logic

    return 0
}