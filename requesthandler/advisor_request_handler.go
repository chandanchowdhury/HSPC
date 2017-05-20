package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/advisor"
	"github.com/go-openapi/runtime/middleware"
)

func HandleAdvisorPost(params advisor.PostAdvisorParams) middleware.Responder {
	//create the advisor
	advisor_id := dbhandler.AdvisorCreate(*params.Advisor)

	if advisor_id <= 0 {
		resp := advisor.NewPostAdvisorDefault(400)
		error := new(models.Error)
		error.Code = advisor_id

		switch advisor_id {
		case -1:
			error.Message = "Duplicate error"
		case -2:
			error.Message = "Related data error"
		default:
			error.Message = "Error: Unexpected error"
		}

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := advisor.NewPostAdvisorOK()
	advisor := dbhandler.AdvisorRead(advisor_id)

	//set response data
	resp.SetPayload(&advisor)

	//return the response
	return resp
}

func HandleAdvisorGet(params advisor.GetAdvisorIDParams) middleware.Responder {
	//get advisor details based on the provided id
	advisor_data := dbhandler.AdvisorRead(params.ID)

	if advisor_data.AdvisorID == 0 {
		resp := advisor.NewGetAdvisorIDDefault(404)
		error := &models.Error{Code: -1, Message: "Advisor not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := advisor.NewGetAdvisorIDOK()
	resp.SetPayload(&advisor_data)

	return resp
}

func HandleAdvisorPut(params advisor.PutAdvisorParams) middleware.Responder {
	affected_count := dbhandler.AdvisorUpdate(*params.Advisor)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := advisor.NewPostAdvisorDefault(400)
		resp.SetPayload(error)

		switch affected_count {
		case 0:
			error.Message = "Warn: no records found for update"
			resp.SetStatusCode(404)
		case -1:
			error.Message = "Update will cause duplicate record"
		case -2:
			error.Message = "Related data error"
		default:
			error.Message = "Error: Unexpected error"
		}

		return resp
	}

	resp := advisor.NewPutAdvisorOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	//return the response
	return resp

}

func HandleAdvisorDelete(params advisor.DeleteAdvisorIDParams) middleware.Responder {
	affected_count := dbhandler.AdvisorDelete(params.ID)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := advisor.NewDeleteAdvisorIDDefault(400)
		resp.SetPayload(error)

		switch affected_count {
		case 0:
			error.Message = "Warn: no records found for delete"
			resp.SetStatusCode(404)
		case -2:
			error.Message = "Related data error"
		default:
			error.Message = "Error: Unexpected error"
		}

		return resp
	}

	resp := advisor.NewDeleteAdvisorIDOK()
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}

func HandleAdvisorReadAll(params advisor.GetAdvisorParams) middleware.Responder {
	advisors := dbhandler.AdvisorReadAll()

	resp := advisor.NewGetAdvisorOK()
	resp.SetPayload(advisors)

	return resp
}

func HandleAdvisorListSchool(params advisor.GetAdvisorIDSchoolsParams, principal interface{}) middleware.Responder {
	// Get School list for advisor
	school_ids := dbhandler.AdvisorGetAllSchools(params.ID)

	schools := make([]*models.School, 0)
	for _, sid := range school_ids {
		school := dbhandler.SchoolRead(sid)
		schools = append(schools, &school)
	}

	resp := advisor.NewGetAdvisorIDSchoolsOK()
	resp.SetPayload(schools)
	return resp
}
