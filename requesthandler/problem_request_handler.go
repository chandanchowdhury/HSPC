package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/problem"
	"github.com/go-openapi/runtime/middleware"
)

func HandleProblemPost(params problem.PostProblemParams) middleware.Responder {
	//create the problem
	error := dbhandler.ProblemCreate(*params.Problem)

	if error.Code != *params.Problem.ProblemID {
		resp := problem.NewPostProblemDefault(400)
		resp.SetPayload(&error)
		return resp
	}

	// create the response
	resp := problem.NewPostProblemOK()
	problem := dbhandler.ProblemRead(*params.Problem.ProblemID)

	//set response data
	resp.SetPayload(&problem)

	//return the response
	return resp
}

func HandleProblemGet(params problem.GetProblemIDParams) middleware.Responder {
	//get problem details based on the provided id
	problem_data := dbhandler.ProblemRead(params.ID)

	if problem_data.ProblemID == nil {
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
	code := dbhandler.ProblemUpdate(*params.Problem)

	error := new(models.Error)

	if code <= 0 {
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
	code := dbhandler.ProblemDelete(params.ID)

	error := new(models.Error)

	if code <= 0 {
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

func HandleProblemGetList(params problem.GetProblemParams) middleware.Responder {
	//get problem list
	problem_list := dbhandler.ProblemReadList()

	if problem_list == nil {
		resp := problem.NewGetProblemIDDefault(404)
		error := &models.Error{Code: -1, Message: "Problem not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := problem.NewGetProblemOK()
	resp.SetPayload(problem_list)

	return resp
}

func HandleGetProblemIDSolutions(params problem.GetProblemIDSolutionsParams) middleware.Responder {
	solution_list := dbhandler.SolutionForProblem(params.ID)

	if solution_list == nil {
		resp := problem.NewGetProblemIDSolutionsDefault(404)
		error := &models.Error{Code: -1, Message: "No Solution found"}

		resp.SetPayload(error)
		return resp
	}

	resp := problem.NewGetProblemIDSolutionsOK()

	resp.SetPayload(solution_list)
	return resp
}
