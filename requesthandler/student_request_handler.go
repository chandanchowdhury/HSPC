package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/student"
	"github.com/go-openapi/runtime/middleware"
)

func HandleStudentPost(params student.PostStudentParams) middleware.Responder {
	//create the student
	student_id := dbhandler.StudentCreate(*params.Student)

	if student_id < 0 {
		resp := student.NewPostStudentDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create Student"

		if student_id == -2 {
			error.Message = "Related data error"
			error.Fields = "School"
		}

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := student.NewPostStudentOK()
	student := dbhandler.StudentRead(student_id)

	//set response data
	resp.SetPayload(&student)

	//return the response
	return resp
}

func HandleStudentGet(params student.GetStudentIDParams) middleware.Responder {
	//get student details based on the provided id
	student_data := dbhandler.StudentRead(params.ID)

	if student_data.StudentID == 0 {
		resp := student.NewGetStudentIDDefault(404)
		error := &models.Error{Code: -1, Message: "Student not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := student.NewGetStudentIDOK()
	resp.SetPayload(&student_data)

	return resp
}

func HandleStudentPut(params student.PutStudentParams) middleware.Responder {
	affected_count := dbhandler.StudentUpdate(*params.Student)

	error := new(models.Error)

	if affected_count < 0 {
		error.Message = "Error: Unexpected number of updates"

		if affected_count == -2 {
			error.Message = "Related data error"
		}

		error.Code = affected_count
		resp := student.NewPostStudentDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := student.NewPutStudentOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleStudentDelete(params student.DeleteStudentIDParams) middleware.Responder {
	affected_count := dbhandler.StudentDelete(params.ID)

	error := new(models.Error)

	if affected_count < 0 {
		error.Message = "Error: Unexpected number of deletes"

		if affected_count == -2 {
			error.Message = "Related data error"
		}

		error.Code = affected_count
		resp := student.NewDeleteStudentIDDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := student.NewDeleteStudentIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
