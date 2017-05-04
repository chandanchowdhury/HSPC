package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/problem"
	"github.com/go-openapi/runtime/middleware"
)

func HandleProblemPost(params problem.PostProblemParams) middleware.Responder {
	//create the problem
	problem_id := dbhandler.ProblemCreate(*params.Problem)

	if problem_id == -1 {
		resp := problem.NewPostProblemDefault(400)
		error := new(models.Error)
		error.Code = -1
		error.Message = "Failed to create Problem"

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := problem.NewPostProblemOK()
	error := new(models.Error)

	error.Code = problem_id
	error.Message = "Created"

	//set response data
	resp.SetPayload(error)

	//return the response
	return resp
}

func HandleProblemGet(params problem.GetProblemIDParams) middleware.Responder {
	//get problem details based on the provided id
	problem_data := dbhandler.ProblemRead(params.ID)

	if problem_data.ProblemID == 0 {
		resp := problem.NewGetProblemIDDefault(404)
		error := &models.Error{Code: -1, Message: "Problem not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := problem.NewGetProblemIDOK()
	resp.SetPayload(&problem_data)

	return resp
}

func HandleProblemPut(params problem.PutProblemParams) middleware.Responder {
	success := dbhandler.ProblemUpdate(*params.Problem)

	error := new(models.Error)

	if !success {
		error.Message = "Error: update failed"
		resp := problem.NewPostProblemDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := problem.NewPutProblemOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleProblemDelete(params problem.DeleteProblemIDParams) middleware.Responder {
	success := dbhandler.ProblemDelete(params.ID)

	error := new(models.Error)

	if !success {
		error.Message = "Error: Unexpected number of deletes"
		resp := problem.NewDeleteProblemIDDefault(400)
		resp.SetPayload(error)

		return resp
	}

	resp := problem.NewDeleteProblemIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
