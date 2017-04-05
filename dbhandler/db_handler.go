package dbhandler

import (
	"database/sql"
	"fmt"
	// the driver is used internally, the underscore makes sure the "unused"
	// error is suppressed.
	_ "github.com/lib/pq"
	"log"
	"github.com/chandanchowdhury/HSPC/models"
)

//TODO: Read from a config file or environment
const (
	DB_USER     = "hspc"
	DB_PASSWORD = "HSPC-Password"
	DB_NAME     = "postgres"
)

func getDBConn() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Print("Error connecting DB")
		log.Fatal(err)
	}
	return db
}

/**
Credential
 **/
func CredentialCreate(credential models.Credential) int64 {
	log.Print("# Creating credential")

	db := getDBConn()

	stmt, err := db.Prepare("INSERT INTO Credential(emailaddress, password_hash)" +
		" VALUES($1, $2) returning credential_id")
	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	var lastInsertId int64
	err = stmt.QueryRow(credential.Emailaddress, credential.Password).Scan(&lastInsertId)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("credential_id = %d", lastInsertId)

	return lastInsertId
}

func CredentialRead(emailaddress string) models.Credential {
	log.Print("# Reading Credential")
	log.Printf("emailaddress = %s", emailaddress)

	db := getDBConn()
	stmt, err := db.Prepare("SELECT credential_id, emailaddress, password_hash " +
		"FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	var credential = models.Credential{}
	err = stmt.QueryRow(emailaddress).Scan(
		&credential.CredentialID, &credential.Emailaddress,
		&credential.Password)

	if err == sql.ErrNoRows {
		return models.Credential{}
	}

	if err != nil {
		log.Print("Error reading Credential data")
		log.Fatal(err)
	}

	return credential
}

func CredentialUpdate(emailaddress string, password string) int64 {
	db := getDBConn()

	log.Print("# Updating Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("UPDATE Credential SET emailaddress = $1, " +
		"password_hash = $2 WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	result, err := stmt.Exec(emailaddress, password)

	if err != nil {
		log.Print("Error updating Credential")
		log.Fatal(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func CredentialDelete(emailaddress string) int64 {
	db := getDBConn()

	log.Print("# Deleting Credential")
	log.Printf("emailaddress = %s", emailaddress)

	stmt, err := db.Prepare("DELETE FROM Credential WHERE emailaddress = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	result, err := stmt.Exec(emailaddress)

	if err != nil {
		log.Print("Error deleting from Credential")
		log.Fatal(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

/*
Address
*/
func AddressCreate(address models.Address) int64 {
	db := getDBConn()

	log.Print("# Creating Address")

	stmt, err := db.Prepare("INSERT INTO address(address_country, address_zip, " +
		"address_state, address_city, address_line1, address_line2) " +
		"VALUES($1, $2, $3, $4, $5, $6) returning address_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	var address_id int64
	err = stmt.QueryRow(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2).Scan(&address_id)

	if err != nil {
		log.Fatal(err)
	}

	return address_id
}

func AddressRead(address_id int64) models.Address {
	log.Print("# Reading Address")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT address_id, address_country, address_zip, " +
		"address_city, address_line1, address_line2 " +
		"FROM address WHERE address_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	var address = models.Address{}
	err = stmt.QueryRow(address_id).Scan(&address.AddressID, &address.Country,
		&address.Zipcode, &address.City, &address.Line1, &address.Line2)

	if err == sql.ErrNoRows {
		return models.Address{}
	}

	if err != nil {
		log.Fatal("Error reading address data")
		log.Panic(err)
	}

	return address
}

func AddressUpdate(address models.Address) int64 {
	db := getDBConn()

	log.Print("# Updating Address")
	log.Printf("Address ID = %d", address.AddressID)

	stmt, err := db.Prepare("UPDATE address SET address_country = $1 " +
		" ,address_zip = $2 " +
		" ,address_state = $3 ,address_city = $4" +
		" ,address_line1 = $5 ,address_line2 = $6" +
		" WHERE address_id = $7")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	result, err := stmt.Exec(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2, address.AddressID)

	if err != nil {
		log.Fatal("Error updating")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func AddressDelete(address_id int64) int64 {
	db := getDBConn()

	log.Print("# Deleting Address")
	log.Printf("Address ID = %d", address_id)

	stmt, err := db.Prepare("DELETE FROM address WHERE address_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Fatal(err)
	}

	result, err := stmt.Exec(address_id)

	if err != nil {
		log.Fatal("Delete Failed")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

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
		log.Fatal(err)
	}

	var school_id int64
	err = stmt.QueryRow(school.SchoolName, school.AddressID).Scan(&school_id)

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	result, err := stmt.Exec(school.SchoolName, school.AddressID, school.SchoolID)

	if err != nil {
		log.Fatal("Error updating school")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
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
		log.Fatal(err)
	}

	result, err := stmt.Exec(school_id)

	if err != nil {
		log.Fatal("Delete Failed")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Fatalf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

/*
Advisor
*/
func AdvisorCreate(advisor models.Advisor) int64 {
	//TODO: complete the logic

	return 0
}

func AdvisorRead(advisor_id int64) models.Advisor {
	var advisor = models.Advisor{}

	//TODO: complete the logic

	return advisor
}

func AdvisorUpdate(advisor models.Advisor) models.Advisor {
	//TODO: complete the logic

	return advisor
}

func AdvisorDelete(advisor_id int64) int64 {
	//TODO: complete the logic

	return 0
}

/*
Team
*/
func TeamCreate(team models.Team) int64 {
	//TODO: complete the logic

	return 0
}

func TeamRead(team_id int64) models.Team {
	var team = models.Team{}

	//TODO: complete the logic

	return team
}

func TeamUpdate(team models.Team) models.Team {
	//TODO: complete the logic

	return team
}

func TeamDelete(team_id int64) int64 {
	//TODO: complete the logic

	return 0
}

/*
Student
*/
func studentCreate(student models.Student) int64 {
	//TODO: complete the logic

	return 0
}

func studentRead(student_id int64) models.Student {
	var student = models.Student{}

	//TODO: complete the logic

	return student
}

func studentUpdate(student models.Student) models.Student {
	//TODO: complete the logic

	return student
}

func studentDelete(student_id int64) int64 {
	//TODO: complete the logic

	return 0
}

/*
//Team Score

func teamscoreCreate(teamscore team_score_struct) int64 {
	//TODO: complete the logic

	return 0
}

func teamscoreRead(teamscore_id int64) team_score_struct {
	var teamscore = team_score_struct{}

	//TODO: complete the logic

	return teamscore
}

func teamscoreUpdate(teamscore team_score_struct) team_score_struct {
	//TODO: complete the logic

	return teamscore
}

func teamscoreDelete(teamscore_id int64) int64 {
	//TODO: complete the logic

	return 0
}


//Parking

func parkingCreate(parking team_score_struct) int64 {
	//TODO: complete the logic

	return 0
}

func parkingRead(parking_id int64) team_score_struct {
	var parking = team_score_struct{}

	//TODO: complete the logic

	return parking
}

func parkingUpdate(parking team_score_struct) team_score_struct {
	//TODO: complete the logic

	return parking
}

func parkingDelete(parking_id int64) int64 {
	//TODO: complete the logic

	return 0
}


//Problem

func problemCreate(problem problem_struct) int64 {
	//TODO: complete the logic

	return 0
}

func problemRead(problem_id int64) problem_struct {
	var problem = problem_struct{}

	//TODO: complete the logic

	return problem
}

func problemUpdate(problem problem_struct) problem_struct {
	//TODO: complete the logic

	return problem
}

func problemDelete(problem_id int64) int64 {
	//TODO: complete the logic

	return 0
}

//Solution

func solutionCreate(solution solution_struct) int64 {
	//TODO: complete the logic

	return 0
}

func solutionRead(solution_id int64) solution_struct {
	var solution = solution_struct{}

	//TODO: complete the logic

	return solution
}

func solutionUpdate(solution solution_struct) solution_struct {
	//TODO: complete the logic

	return solution
}

func solutionDelete(solution_id int64) int64 {
	//TODO: complete the logic

	return 0
}

//Problem_Solution

func problemsolutionCreate(problemsolution problem_solution_struct) int64 {
	//TODO: complete the logic

	return 0
}

func problemsolutionRead(problemsolution_id int64) problem_solution_struct {
	var problemsolution = problem_solution_struct{}

	//TODO: complete the logic

	return problemsolution
}

func problemsolutionUpdate(problemsolution problem_solution_struct) problem_solution_struct {
	//TODO: complete the logic

	return problemsolution
}

func problemsolutionDelete(problemsolution_id int64) int64 {
	//TODO: complete the logic

	return 0
}
*/
