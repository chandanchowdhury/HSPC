package requesthandler

import (
	"log"

	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
	"github.com/go-openapi/runtime/middleware"
)

/**
Check if the Advisor has access to the School
*/
func checkAdvisorAccessSchool(principal *models.Principal, school_id int64) bool {
	//Only School advisor or admin should get the list of Teams
	log.Printf("Checking Advisor = %s access for School = %d", principal.Email, school_id)

	// Get the Advisor detail using the principal
	advisor := dbhandler.AdvisorReadByEmail(principal.Email)

	//get the school advisor
	school_advisor := dbhandler.SchoolReadAdvisor(school_id)

	if advisor.AdvisorID == school_advisor {
		log.Print("Access Allowed")
		return true
	}

	log.Print("Access NOT Allowed")
	return false
}

func HandleSchoolPost(params school.PostSchoolParams, principal *models.Principal) middleware.Responder {
	//create the school
	school_id := dbhandler.SchoolCreate(*params.School)

	if school_id <= 0 {
		resp := school.NewPostSchoolDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create School"

		switch school_id {
		case -1:
			error.Message = "School already exists"
		case -2:
			error.Message = "Related data error"
		default:
			error.Message = "Error: Unexpected error"
		}

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

func HandleSchoolGet(params school.GetSchoolIDParams, principal *models.Principal) middleware.Responder {
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

func HandleSchoolPut(params school.PutSchoolParams, principal *models.Principal) middleware.Responder {
	affected_count := dbhandler.SchoolUpdate(*params.School)

	error := new(models.Error)

	if affected_count <= 0 {
		error.Message = "Error: Unexpected error"
		error.Code = affected_count
		resp := school.NewPostSchoolDefault(400)

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

		resp.SetPayload(error)

		return resp
	}

	resp := school.NewPutSchoolOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	//return the response
	return resp

}

func HandleSchoolDelete(params school.DeleteSchoolIDParams, principal *models.Principal) middleware.Responder {
	affected_count := dbhandler.SchoolDelete(params.ID)

	error := new(models.Error)

	if affected_count <= 0 {
		error.Message = "Error: Unexpected error"
		error.Code = affected_count
		resp := school.NewDeleteSchoolIDDefault(400)

		switch affected_count {
		case 0:
			error.Message = "Warn: no records found for update"
			resp.SetStatusCode(404)
		case -2:
			error.Message = "Related data error"
		default:
			error.Message = "Error: Unexpected error"
		}

		resp.SetPayload(error)

		return resp
	}

	resp := school.NewDeleteSchoolIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}

func HandleSchoolGetList(params school.GetSchoolParams, principal *models.Principal) middleware.Responder {
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

func HandleSchoolGetStudentList(params school.GetSchoolIDStudentsParams, principal *models.Principal) middleware.Responder {
	log.Printf("SchoolStudentList principal = %s", principal.Email)

	//Only School advisor
	if result := checkAdvisorAccessSchool(principal, params.ID); result == false {
		error := new(models.Error)
		error.Code = 403
		error.Message = "Error: Advisor Not Approved For School"
		resp := school.NewGetSchoolIDStudentsDefault(403)
		resp.SetPayload(error)

		return resp
	}

	student_list := dbhandler.StudentListBySchool(params.ID)

	if student_list == nil {
		error := new(models.Error)
		error.Code = 500
		error.Message = "Error: Failed to get school list from DB"
		resp := school.NewGetSchoolIDStudentsDefault(500)
		resp.SetPayload(error)

		return resp
	}

	resp := school.NewGetSchoolIDStudentsOK()
	resp.SetPayload(student_list)

	return resp
}

func HandleSchoolGetTeamList(params school.GetSchoolIDTeamsParams, principal *models.Principal) middleware.Responder {
	log.Printf("SchoolTeamList SchooldID = %d", params.ID)

	//Only School advisor
	if result := checkAdvisorAccessSchool(principal, params.ID); result == false {
		error := new(models.Error)
		error.Code = 403
		error.Message = "Error: Advisor Not Approved For School"
		resp := school.NewGetSchoolIDTeamsDefault(403)
		resp.SetPayload(error)

		return resp
	}

	//Get the list of team_ids
	team_id_list := dbhandler.TeamListForSchool(params.ID)

	if team_id_list == nil {
		error := new(models.Error)
		error.Code = 500
		error.Message = "Error: Failed to get school list from DB"
		resp := school.NewGetSchoolIDTeamsDefault(500)
		resp.SetPayload(error)

		return resp
	}

	//from the list of team_ids, create the list of Teams
	teams := make([]*models.Team, 0)

	for _, tid := range team_id_list {
		team := dbhandler.TeamRead(tid)
		teams = append(teams, &team)
	}

	resp := school.NewGetSchoolIDTeamsOK()
	resp.SetPayload(teams)

	return resp
}

/**
Add an Advisor to a School
*/
func SchoolAddAdvisor(params school.PutSchoolSchoolIDAdvisorAdvisorIDParams, principal *models.Principal) middleware.Responder {
	log.Printf("Set Advisor = %d for School = %d", params.AdvisorID, params.SchoolID)

	//try to set the advisor for the school
	if dbhandler.SchoolAddAdvisor(params.SchoolID, params.AdvisorID) {
		resp := school.NewPutSchoolSchoolIDAdvisorAdvisorIDNoContent()
		return resp
	}

	resp := school.NewPutSchoolSchoolIDAdvisorAdvisorIDDefault(400)
	return resp
}

/**
Remove an Advisor to a School
*/
func SchoolRemoveAdvisor(params school.DeleteSchoolSchoolIDAdvisorAdvisorIDParams, principal *models.Principal) middleware.Responder {
	log.Printf("Remove Advisor = %d for School = %d", params.AdvisorID, params.SchoolID)

	//try to set the advisor for the school
	if dbhandler.SchoolDeleteAdvisor(params.SchoolID, params.AdvisorID) {
		resp := school.NewDeleteSchoolSchoolIDAdvisorAdvisorIDNoContent()
		return resp
	}

	resp := school.NewDeleteSchoolSchoolIDAdvisorAdvisorIDDefault(400)
	return resp
}
