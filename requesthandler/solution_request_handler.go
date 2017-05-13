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

	if solution_id <= 0 {
		resp := solution.NewPostSolutionDefault(400)
		error := new(models.Error)
		error.Code = solution_id
		error.Message = "Failed to create Solution"

		if solution_id == -2 {
			error.Message = "ProblemID does not exists"
		}

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := solution.NewPostSolutionOK()
	solution := dbhandler.SolutionRead(solution_id)

	//set response data
	resp.SetPayload(&solution)

	//return the response
	return resp
}

func HandleSolutionGet(params solution.GetSolutionIDParams) middleware.Responder {
	//get solution details based on the provided id
	solution_data := dbhandler.SolutionRead(params.ID)

	if solution_data.SolutionID == nil {
		resp := solution.NewGetSolutionIDDefault(404)
		error := &models.Error{}
		error.Message = "Solution not found"

		resp.SetPayload(error)
		return resp
	}

	resp := solution.NewGetSolutionIDOK()
	resp.SetPayload(&solution_data)

	return resp
}

func HandleSolutionPut(params solution.PutSolutionParams) middleware.Responder {
	code := dbhandler.SolutionUpdate(*params.Solution)

	error := new(models.Error)

	if code <= 0 {
		error.Message = "Error: Unexpected number of updates"
		resp := solution.NewPostSolutionDefault(400)

		if code == 0 {
			error.Message = "Warn: Solution does not exists"
			resp.SetStatusCode(404)
		}

		resp.SetPayload(error)

		return resp
	}

	resp := solution.NewPutSolutionOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleSolutionDelete(params solution.DeleteSolutionIDParams) middleware.Responder {
	code := dbhandler.SolutionDelete(params.ID)

	error := new(models.Error)

	if code <= 0 {
		error.Message = "Error: Unexpected number of deletes"
		resp := solution.NewDeleteSolutionIDDefault(400)

		if code == 0 {
			error.Message = "Warn: Solution does not exists"
			resp.SetStatusCode(404)
		}
		resp.SetPayload(error)

		return resp
	}

	resp := solution.NewDeleteSolutionIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
