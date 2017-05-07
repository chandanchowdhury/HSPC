package dbhandler

import (
	"database/sql"
	"github.com/chandanchowdhury/HSPC/models"
	"log"
)

/*
Student
*/
func StudentCreate(student models.Student) int64 {
	log.Print("# Creating Student")

	db := getDBConn()
	stmt, err := db.Prepare("INSERT INTO student(student_name, student_grade, school_id) " +
		"VALUES($1, $2, $3) returning student_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	var student_id int64
	err = stmt.QueryRow(student.StudentName, student.StudentGrade, student.SchoolID).Scan(&student_id)

	if err != nil {
		log.Print(err)
		return 0
	}

	return student_id
}

func StudentRead(student_id int64) models.Student {
	var student = models.Student{}

	log.Print("# Reading Student")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT student_id, student_name, student_grade, school_id " +
		"FROM student WHERE student_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

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
	db := getDBConn()

	log.Print("# Updating Student")
	log.Printf("Student ID = %d", student.StudentID)

	stmt, err := db.Prepare("UPDATE student SET student_name = $1, student_grade = $2, school_id = $3" +
		"WHERE student_id = $4")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(student.StudentName, student.StudentGrade, student.SchoolID, student.StudentID)

	if err != nil {
		log.Print("Error updating student")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Printf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func StudentDelete(student_id int64) int64 {
	db := getDBConn()

	log.Print("# Deleting Student")
	log.Printf("Student ID = %d", student_id)

	stmt, err := db.Prepare("DELETE FROM student WHERE student_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(student_id)

	if err != nil {
		log.Print("Delete Failed")
		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Printf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

/**
	Get all Student who belongs to a School
 */
func StudentListBySchool(school_id int64) []*models.Student {
	log.Print("# Reading Student List for School")

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
