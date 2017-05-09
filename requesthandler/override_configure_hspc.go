package requesthandler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/chandanchowdhury/HSPC/restapi/operations"
	"github.com/chandanchowdhury/HSPC/restapi/operations/address"
	"github.com/chandanchowdhury/HSPC/restapi/operations/advisor"
	"github.com/chandanchowdhury/HSPC/restapi/operations/credential"
	"github.com/chandanchowdhury/HSPC/restapi/operations/login"
	"github.com/chandanchowdhury/HSPC/restapi/operations/problem"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
	"github.com/chandanchowdhury/HSPC/restapi/operations/solution"
	"github.com/chandanchowdhury/HSPC/restapi/operations/student"
	"github.com/chandanchowdhury/HSPC/restapi/operations/team"
)

/**
This function when called inside configureAPI() function in
 configure_hspc.go will override all NotImplemented calls with actual
 valid calls to handle the request.

Just add below line before return in configureAPI()

requesthandler.Override_configure_hspc(api)
*/
func Override_configure_hspc(api *operations.HspcAPI) {

	api.AddressDeleteAddressIDHandler = address.DeleteAddressIDHandlerFunc(func(params address.DeleteAddressIDParams) middleware.Responder {
		return HandleAddressDelete(params)
	})
	api.AdvisorDeleteAdvisorIDHandler = advisor.DeleteAdvisorIDHandlerFunc(func(params advisor.DeleteAdvisorIDParams) middleware.Responder {
		return HandleAdvisorDelete(params)
	})
	api.CredentialDeleteCredentialIDHandler = credential.DeleteCredentialIDHandlerFunc(func(params credential.DeleteCredentialIDParams) middleware.Responder {
		return HandleCredentialDelete(params)
	})
	api.ProblemDeleteProblemIDHandler = problem.DeleteProblemIDHandlerFunc(func(params problem.DeleteProblemIDParams) middleware.Responder {
		return HandleProblemDelete(params)
	})
	api.SchoolDeleteSchoolIDHandler = school.DeleteSchoolIDHandlerFunc(func(params school.DeleteSchoolIDParams) middleware.Responder {
		return HandleSchoolDelete(params)
	})
	api.SolutionDeleteSolutionIDHandler = solution.DeleteSolutionIDHandlerFunc(func(params solution.DeleteSolutionIDParams) middleware.Responder {
		return HandleSolutionDelete(params)
	})
	api.StudentDeleteStudentIDHandler = student.DeleteStudentIDHandlerFunc(func(params student.DeleteStudentIDParams) middleware.Responder {
		return HandleStudentDelete(params)
	})
	api.TeamDeleteTeamIDHandler = team.DeleteTeamIDHandlerFunc(func(params team.DeleteTeamIDParams) middleware.Responder {
		return HandleTeamDelete(params)
	})
	api.AddressGetAddressIDHandler = address.GetAddressIDHandlerFunc(func(params address.GetAddressIDParams) middleware.Responder {
		return HandleAddressGet(params)
	})
	api.AdvisorGetAdvisorIDHandler = advisor.GetAdvisorIDHandlerFunc(func(params advisor.GetAdvisorIDParams) middleware.Responder {
		return HandleAdvisorGet(params)
	})
	api.CredentialGetCredentialIDHandler = credential.GetCredentialIDHandlerFunc(func(params credential.GetCredentialIDParams) middleware.Responder {
		return HandleCredentialGet(params)
	})
	api.LoginGetLoginEmailaddressPasswordHandler = login.GetLoginEmailaddressPasswordHandlerFunc(func(params login.GetLoginEmailaddressPasswordParams) middleware.Responder {
		return HandleLogin(params)
	})
	api.ProblemGetProblemIDHandler = problem.GetProblemIDHandlerFunc(func(params problem.GetProblemIDParams) middleware.Responder {
		return HandleProblemGet(params)
	})
	api.SchoolGetSchoolIDHandler = school.GetSchoolIDHandlerFunc(func(params school.GetSchoolIDParams) middleware.Responder {
		return HandleSchoolGet(params)
	})
	api.SolutionGetSolutionIDHandler = solution.GetSolutionIDHandlerFunc(func(params solution.GetSolutionIDParams) middleware.Responder {
		return HandleSolutionGet(params)
	})
	api.StudentGetStudentIDHandler = student.GetStudentIDHandlerFunc(func(params student.GetStudentIDParams) middleware.Responder {
		return HandleStudentGet(params)
	})
	api.TeamGetTeamIDHandler = team.GetTeamIDHandlerFunc(func(params team.GetTeamIDParams) middleware.Responder {
		return HandleTeamGet(params)
	})
	api.AddressPostAddressHandler = address.PostAddressHandlerFunc(func(params address.PostAddressParams) middleware.Responder {
		return HandleAddressPost(params)
	})
	api.AdvisorPostAdvisorHandler = advisor.PostAdvisorHandlerFunc(func(params advisor.PostAdvisorParams) middleware.Responder {
		return HandleAdvisorPost(params)
	})
	api.CredentialPostCredentialHandler = credential.PostCredentialHandlerFunc(func(params credential.PostCredentialParams) middleware.Responder {
		return HandleCredentialPost(params)
	})
	api.ProblemPostProblemHandler = problem.PostProblemHandlerFunc(func(params problem.PostProblemParams) middleware.Responder {
		return HandleProblemPost(params)
	})
	api.SchoolPostSchoolHandler = school.PostSchoolHandlerFunc(func(params school.PostSchoolParams) middleware.Responder {
		return HandleSchoolPost(params)
	})
	api.SolutionPostSolutionHandler = solution.PostSolutionHandlerFunc(func(params solution.PostSolutionParams) middleware.Responder {
		return HandleSolutionPost(params)
	})
	api.StudentPostStudentHandler = student.PostStudentHandlerFunc(func(params student.PostStudentParams) middleware.Responder {
		return HandleStudentPost(params)
	})
	api.TeamPostTeamHandler = team.PostTeamHandlerFunc(func(params team.PostTeamParams) middleware.Responder {
		return HandleTeamPost(params)
	})
	api.AddressPutAddressHandler = address.PutAddressHandlerFunc(func(params address.PutAddressParams) middleware.Responder {
		return HandleAddressPut(params)
	})
	api.AdvisorPutAdvisorHandler = advisor.PutAdvisorHandlerFunc(func(params advisor.PutAdvisorParams) middleware.Responder {
		return HandleAdvisorPut(params)
	})
	api.CredentialPutCredentialHandler = credential.PutCredentialHandlerFunc(func(params credential.PutCredentialParams) middleware.Responder {
		return HandleCredentialPut(params)
	})
	api.ProblemPutProblemHandler = problem.PutProblemHandlerFunc(func(params problem.PutProblemParams) middleware.Responder {
		return HandleProblemPut(params)
	})
	api.SchoolPutSchoolHandler = school.PutSchoolHandlerFunc(func(params school.PutSchoolParams) middleware.Responder {
		return HandleSchoolPut(params)
	})
	api.SolutionPutSolutionHandler = solution.PutSolutionHandlerFunc(func(params solution.PutSolutionParams) middleware.Responder {
		return HandleSolutionPut(params)
	})
	api.StudentPutStudentHandler = student.PutStudentHandlerFunc(func(params student.PutStudentParams) middleware.Responder {
		return HandleStudentPut(params)
	})
	api.TeamPutTeamHandler = team.PutTeamHandlerFunc(func(params team.PutTeamParams) middleware.Responder {
		return HandleTeamPut(params)
	})
	// List of all Schools
	api.SchoolGetSchoolHandler = school.GetSchoolHandlerFunc(func(params school.GetSchoolParams) middleware.Responder {
		return HandleSchoolGetList(params)
	})

	// List of all Problems
	api.ProblemGetProblemHandler = problem.GetProblemHandlerFunc(func(params problem.GetProblemParams) middleware.Responder {
		return HandleProblemGetList(params)
	})
	// List of all Students for a School
	api.SchoolGetSchoolIDStudentsHandler = school.GetSchoolIDStudentsHandlerFunc(func(params school.GetSchoolIDStudentsParams) middleware.Responder {
		return HandleSchoolGetStudentList(params)
	})
}
