package dbhandler

import (
	"database/sql"
	"log"

	"github.com/chandanchowdhury/HSPC/models"
)

/**
Create a Student

*/
func StudentCreate(student models.Student) int64 {
	log.Print("Creating Student")

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO student(student_name, student_grade, school_id) " +
		"VALUES($1, $2, $3) returning student_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var student_id int64
	err = stmt.QueryRow(student.StudentName, student.StudentGrade, student.SchoolID).Scan(&student_id)

	if err != nil {
		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	log.Printf("New student_id = %d", student_id)

	return student_id
}

func StudentRead(student_id int64) models.Student {
	log.Printf("Reading Student ID = %d", student_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT student_id, student_name, student_grade, school_id " +
		"FROM student WHERE student_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var student = models.Student{}
	err = stmt.QueryRow(student_id).Scan(&student.StudentID, &student.StudentName, &student.StudentGrade, &student.SchoolID)

	// if no records found, return an empty struct
	if err == sql.ErrNoRows {
		return models.Student{}
	}

	if err != nil {
		log.Print("Error getting student data")
		log.Panic(err)
	}

	return student
}

func StudentUpdate(student models.Student) int64 {
	log.Printf("Updating Student ID = %d", student.StudentID)

	db := getDBConn()

	stmt, err := db.Prepare("UPDATE student SET student_name = $1, student_grade = $2, school_id = $3" +
		"WHERE student_id = $4")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := stmt.Exec(student.StudentName, student.StudentGrade, student.SchoolID, student.StudentID)

	if err != nil {
		log.Print("Error updating student")

		if isDuplicateKeyError(err) {
			return -1
		}

		if isForeignKeyError(err) {
			return -2
		}

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

func StudentDelete(student_id int64) int64 {
	log.Printf("Deleting Student ID = %d", student_id)

	db := getDBConn()

	stmt, err := db.Prepare("DELETE FROM student WHERE student_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	result, err := stmt.Exec(student_id)

	if err != nil {
		log.Print("Delete Failed")

		if isForeignKeyError(err) {
			return -2
		}

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

/**
StudentListBySchool Finds and returns all Student who belongs to a School
*/
func StudentListBySchool(school_id int64) []*models.Student {
	log.Printf("Reading Student List for School ID = %d", school_id)

	db := getDBConn()

	stmt, err := db.Prepare("SELECT student_id, student_name, student_grade, school_id " +
		"FROM student WHERE school_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	crsr, err := stmt.Query(school_id)

	// if no records found, return an empty struct
	if err == sql.ErrNoRows {
		return []*models.Student{}
	}

	if err != nil {
		log.Print("Error getting student data")
		log.Panic(err)
	}

	student_list := make([]*models.Student, 0)
	for crsr.Next() {
		student := new(models.Student)
		crsr.Scan(&student.StudentID, &student.StudentName, &student.StudentGrade, &student.SchoolID)

		student_list = append(student_list, student)
	}

	return student_list
}
