package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
	"github.com/go-openapi/runtime/middleware"
)

func HandleSchoolPost(params school.PostSchoolParams) middleware.Responder {
	//create the school
	school_id := dbhandler.SchoolCreate(*params.School)

	if school_id == -1 {
		resp := school.NewPostSchoolDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create School"

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := school.NewPostSchoolOK()
	school := dbhandler.SchoolRead(school_id)

	//set response data
	resp.SetPayload(&school)

	//return the response
	return resp
}

func HandleSchoolGet(params school.GetSchoolIDParams) middleware.Responder {
	//get school details based on the provided id
	school_data := dbhandler.SchoolRead(params.ID)

	if school_data.SchoolID == 0 {
		resp := school.NewGetSchoolIDDefault(404)
		error := &models.Error{Code: -1, Message: "School not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := school.NewGetSchoolIDOK()
	resp.SetPayload(&school_data)

	return resp
}

func HandleSchoolPut(params school.PutSchoolParams) middleware.Responder {
	affected_count := dbhandler.SchoolUpdate(*params.School)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := school.NewPostSchoolDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := school.NewPutSchoolOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	//return the response
	return resp

}

func HandleSchoolDelete(params school.DeleteSchoolIDParams) middleware.Responder {
	affected_count := dbhandler.SchoolDelete(params.ID)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := school.NewDeleteSchoolIDDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := school.NewDeleteSchoolIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}

func HandleSchoolGetList(params school.GetSchoolParams) middleware.Responder {
	school_list := dbhandler.SchoolList(params)

	if school_list == nil {
		error := new(models.Error)
		error.Code = 500
		error.Message = "Error: Failed to get school list from DB"
		resp := school.NewGetSchoolDefault(500)
		resp.SetPayload(error)

		return resp
	}

	resp := school.NewGetSchoolOK()
	resp.SetPayload(school_list)

	return resp
}

func HandleSchoolGetStudentList(params school.GetSchoolIDStudentsParams) middleware.Responder {
	student_list := dbhandler.StudentListBySchool(params.ID)

	if student_list == nil {
		error := new(models.Error)
		error.Code = 500
		error.Message = "Error: Failed to get school list from DB"
		resp := school.NewGetSchoolDefault(500)
		resp.SetPayload(error)

		return resp
	}

	resp := school.NewGetSchoolIDStudentsOK()
	resp.SetPayload(student_list)

	return resp
}
