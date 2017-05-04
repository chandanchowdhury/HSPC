package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/solution"
	"github.com/go-openapi/runtime/middleware"
)

func HandleSolutionPost(params solution.PostSolutionParams) middleware.Responder {
	//create the solution
	solution_id := dbhandler.SolutionCreate(*params.Solution)

	if solution_id == -1 {
		resp := solution.NewPostSolutionDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create Solution"

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := solution.NewPostSolutionOK()
	error := new(models.Error)

	error.Message = "Created"

	//set response data
	resp.SetPayload(error)

	//return the response
	return resp
}

func HandleSolutionGet(params solution.GetSolutionIDParams) middleware.Responder {
	//get solution details based on the provided id
	solution_data := dbhandler.SolutionRead(params.ID)

	if solution_data.SolutionID == 0 {
		resp := solution.NewGetSolutionIDDefault(404)
		error := &models.Error{Code: -1, Message: "Solution not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := solution.NewGetSolutionIDOK()
	resp.SetPayload(&solution_data)

	return resp
}

func HandleSolutionPut(params solution.PutSolutionParams) middleware.Responder {
	success := dbhandler.SolutionUpdate(*params.Solution)

	error := new(models.Error)

	if !success {
		error.Message = "Error: Unexpected number of updates"
		resp := solution.NewPostSolutionDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := solution.NewPutSolutionOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleSolutionDelete(params solution.DeleteSolutionIDParams) middleware.Responder {
	success := dbhandler.SolutionDelete(params.ID)

	error := new(models.Error)

	if !success {
		error.Message = "Error: Unexpected number of deletes"
		resp := solution.NewDeleteSolutionIDDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := solution.NewDeleteSolutionIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
