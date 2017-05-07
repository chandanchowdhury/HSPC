package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/team"
	"github.com/go-openapi/runtime/middleware"
)

func HandleTeamPost(params team.PostTeamParams) middleware.Responder {
	//create the team
	team_id := dbhandler.TeamCreate(*params.Team)

	if team_id == -1 {
		resp := team.NewPostTeamDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create Team"

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

func HandleTeamGet(params team.GetTeamIDParams) middleware.Responder {
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

func HandleTeamPut(params team.PutTeamParams) middleware.Responder {
	affected_count := dbhandler.TeamUpdate(*params.Team)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := team.NewPostTeamDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := team.NewPutTeamOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleTeamDelete(params team.DeleteTeamIDParams) middleware.Responder {
	affected_count := dbhandler.TeamDelete(params.ID)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := team.NewDeleteTeamIDDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := team.NewDeleteTeamIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
