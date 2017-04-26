package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
	"github.com/go-openapi/runtime/middleware"
)

func HandleSchoolPost(params school.PostSchoolParams) middleware.Responder {
	//create the school
	school_id := dbhandler.SchoolCreate(*params.Body)

	school_data := dbhandler.SchoolRead(school_id)

	// create the response
	resp := school.NewGetSchoolIDOK()

	//set response data
	resp.SetPayload(&school_data)

	//return the response
	return resp
}

func HandleSchoolGet(params school.GetSchoolIDParams) middleware.Responder {
	//get school details based on the provided id
	school_data := dbhandler.SchoolRead(params.ID)

	resp := school.NewGetSchoolIDOK()

	resp.SetPayload(&school_data)

	return resp
}

func HandleSchoolPut(params school.PutSchoolParams) middleware.Responder {
	affected_count := dbhandler.SchoolUpdate(*params.School)

	var resp_str string
	resp_str = "Success"

	if affected_count != 1 {
		resp_str = "Error"
	}

	resp := school.NewPutSchoolOK()
	resp.SetPayload(resp_str)

	return resp
}

func HandleSchoolDelete(params school.DeleteSchoolIDParams) middleware.Responder {
	affected_count := dbhandler.SchoolDelete(params.ID)

	var resp_str string
	resp_str = "Success"

	if affected_count != 1 {
		resp_str = "Error"
	}

	resp := school.NewDeleteSchoolIDOK()
	resp.SetPayload(resp_str)

	return resp
}
