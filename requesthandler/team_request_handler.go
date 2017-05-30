package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/team"
	"github.com/go-openapi/runtime/middleware"
)

func checkAdvisorAccessTeam(principal *models.Principal, team_id int64) bool {
	advisor := dbhandler.AdvisorReadByEmail(principal.Email)

	//Get Schools Advisor is approved for
	advisor_school_ids := dbhandler.AdvisorGetAllSchools(advisor.AdvisorID)

	//Get SchoolID for the TeamID
	team_detail := dbhandler.TeamRead(team_id)

	//if SchoolID in list of approved Schools, allow
	for _, sid := range advisor_school_ids {
		if sid == *team_detail.SchoolID {
			return true
		}
	}

	return true
}

func HandleTeamPost(params team.PostTeamParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.Team.TeamID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	//create the team
	team_id := dbhandler.TeamCreate(*params.Team)

	if team_id <= 0 {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = team_id

		switch team_id {
		case 0:
			error.Message = "Warn: no records found for update"
			resp.SetStatusCode(404)
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
	resp := team.NewPostTeamOK()
	team := dbhandler.TeamRead(team_id)

	//set response data
	resp.SetPayload(&team)

	//return the response
	return resp
}

func HandleTeamGet(params team.GetTeamIDParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.ID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	//get team details based on the provided id
	team_data := dbhandler.TeamRead(params.ID)

	if team_data.TeamID == 0 {
		resp := team.NewGetTeamIDDefault(404)
		error := &models.Error{Code: -1, Message: "Team not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := team.NewGetTeamIDOK()
	resp.SetPayload(&team_data)

	return resp
}

func HandleTeamPut(params team.PutTeamParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.Team.TeamID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	affected_count := dbhandler.TeamUpdate(*params.Team)

	error := new(models.Error)

	if affected_count <= 0 {
		error.Code = affected_count
		resp := team.NewPostTeamDefault(400)

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

	resp := team.NewPutTeamOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleTeamDelete(params team.DeleteTeamIDParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.ID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	affected_count := dbhandler.TeamDelete(params.ID)

	error := new(models.Error)

	if affected_count <= 0 {
		error.Code = affected_count
		resp := team.NewDeleteTeamIDDefault(400)

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

	resp := team.NewDeleteTeamIDOK()
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}

func HandleGetTeamStudents(params team.GetTeamIDStudentsParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.ID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	//get the list of Student_ids
	team_member := dbhandler.TeamReadMembers(params.ID)

	//prepare the list of Students
	if team_member != nil && len(team_member) > 0 {
		student_list := make([]*models.Student, 0)
		for _, id := range team_member {
			student := dbhandler.StudentRead(id)

			student_list = append(student_list, &student)
		}

		resp := team.NewGetTeamIDStudentsOK()

		resp.SetPayload(student_list)
		return resp
	}

	error := new(models.Error)
	error.Code = -1
	error.Message = "Unexpected error"
	resp := team.NewGetTeamIDStudentsDefault(400)
	resp.SetPayload(error)
	return resp

}

func HandleTeamAddStudent(params team.GetTeamTeamIDAddStudentIDParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.TeamID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	if success := dbhandler.TeamAddMember(params.TeamID, params.StudentID); success == true {
		// just return 201 without any content
		resp := team.NewGetTeamTeamIDAddStudentIDNoContent()

		return resp
	}

	error := new(models.Error)
	error.Code = -1
	error.Message = "Unexpected error"
	resp := team.NewGetTeamIDStudentsDefault(400)
	resp.SetPayload(error)
	return resp
}

func HandleTeamRemoveStudent(params team.GetTeamTeamIDRemoveStudentIDParams, principal *models.Principal) middleware.Responder {
	//first check advisor access
	if allowed := checkAdvisorAccessTeam(principal, params.TeamID); allowed == false {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -3
		error.Message = "Advisor Not Approved for School"
		resp.SetPayload(error)
		return resp
	}

	if success := dbhandler.TeamDeleteMember(params.TeamID, params.StudentID); success == true {
		// just return 201 without any content
		resp := team.NewGetTeamTeamIDRemoveStudentIDNoContent()

		return resp
	}

	error := new(models.Error)
	error.Code = -1
	error.Message = "Unexpected error"
	resp := team.NewGetTeamIDStudentsDefault(400)
	resp.SetPayload(error)
	return resp
}
