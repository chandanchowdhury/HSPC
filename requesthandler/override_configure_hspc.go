package requesthandler

import (
	"log"

	"github.com/chandanchowdhury/HSPC/dbhandler"
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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

/**
This function when called inside configureAPI() function in
 configure_hspc.go will override all NotImplemented calls with actual
 valid calls to handle the request.

Just add below line before return in configureAPI()

requesthandler.Override_configure_hspc(api)
*/
//
func Override_configure_hspc(api *operations.HspcAPI) {

	// Applies when the Authorization header is set with the Basic scheme
	api.UserSecurityBasicAuth = func(user string, pass string) (interface{}, error) {
		log.Print("Authenticating User = %s", user)

		credential := dbhandler.CredentialRead(user)

		if credential.CredentialID == 0 {
			return nil, errors.Unauthenticated("api.hspc.ksu.edu")
		}

		log.Printf("From DB - email: %s, password: %s Active: %t", credential.Emailaddress.String(), credential.Password.String(), *credential.CredentialActive)

		if credential.Password.String() == pass && *credential.CredentialActive == true {
			log.Print("Password Matched")
			return user, nil
		}

		return nil, errors.Unauthenticated("api.hspc.ksu.edu")
	}

	// Applies when the Authorization header is set with the Basic scheme
	api.AdminSecurityBasicAuth = func(user string, pass string) (interface{}, error) {
		//TODO: improvement required
		log.Printf("Authenticating Admin = %s", user)

		if user != "hspc" {
			return nil, errors.Unauthenticated("admin level access")
		}

		if user == "hspc" && pass == "hspc" {
			return "HSPC", nil
		}

		return nil, errors.Unauthenticated("api.hspc.cs.ksu.edu")
	}

	api.AddressDeleteAddressIDHandler = address.DeleteAddressIDHandlerFunc(func(params address.DeleteAddressIDParams, principal interface{}) middleware.Responder {
		return HandleAddressDelete(params, principal)
	})
	api.AdvisorDeleteAdvisorIDHandler = advisor.DeleteAdvisorIDHandlerFunc(func(params advisor.DeleteAdvisorIDParams, principal interface{}) middleware.Responder {
		return HandleAdvisorDelete(params)
	})
	api.CredentialDeleteCredentialEmailaddressHandler = credential.DeleteCredentialEmailaddressHandlerFunc(func(params credential.DeleteCredentialEmailaddressParams, principal interface{}) middleware.Responder {
		return HandleCredentialDelete(params)
	})
	api.ProblemDeleteProblemIDHandler = problem.DeleteProblemIDHandlerFunc(func(params problem.DeleteProblemIDParams, principal interface{}) middleware.Responder {
		return HandleProblemDelete(params)
	})
	api.SchoolDeleteSchoolIDHandler = school.DeleteSchoolIDHandlerFunc(func(params school.DeleteSchoolIDParams, principal interface{}) middleware.Responder {
		return HandleSchoolDelete(params)
	})
	api.SolutionDeleteSolutionIDHandler = solution.DeleteSolutionIDHandlerFunc(func(params solution.DeleteSolutionIDParams, principal interface{}) middleware.Responder {
		return HandleSolutionDelete(params)
	})
	api.StudentDeleteStudentIDHandler = student.DeleteStudentIDHandlerFunc(func(params student.DeleteStudentIDParams, principal interface{}) middleware.Responder {
		return HandleStudentDelete(params, principal)
	})
	api.TeamDeleteTeamIDHandler = team.DeleteTeamIDHandlerFunc(func(params team.DeleteTeamIDParams, principal interface{}) middleware.Responder {
		return HandleTeamDelete(params, principal)
	})
	api.AddressGetAddressIDHandler = address.GetAddressIDHandlerFunc(func(params address.GetAddressIDParams, principal interface{}) middleware.Responder {
		return HandleAddressGet(params)
	})
	api.AdvisorGetAdvisorIDHandler = advisor.GetAdvisorIDHandlerFunc(func(params advisor.GetAdvisorIDParams, principal interface{}) middleware.Responder {
		return HandleAdvisorGet(params)
	})
	api.CredentialGetCredentialEmailaddressHandler = credential.GetCredentialEmailaddressHandlerFunc(func(params credential.GetCredentialEmailaddressParams, principal interface{}) middleware.Responder {
		return HandleCredentialGet(params)
	})
	api.LoginGetLoginEmailaddressPasswordHandler = login.GetLoginEmailaddressPasswordHandlerFunc(func(params login.GetLoginEmailaddressPasswordParams, principal interface{}) middleware.Responder {
		return HandleLogin(params)
	})
	api.ProblemGetProblemIDHandler = problem.GetProblemIDHandlerFunc(func(params problem.GetProblemIDParams, principal interface{}) middleware.Responder {
		return HandleProblemGet(params)
	})
	api.SchoolGetSchoolIDHandler = school.GetSchoolIDHandlerFunc(func(params school.GetSchoolIDParams, principal interface{}) middleware.Responder {
		return HandleSchoolGet(params)
	})
	api.SolutionGetSolutionIDHandler = solution.GetSolutionIDHandlerFunc(func(params solution.GetSolutionIDParams, principal interface{}) middleware.Responder {
		return HandleSolutionGet(params)
	})
	api.StudentGetStudentIDHandler = student.GetStudentIDHandlerFunc(func(params student.GetStudentIDParams, principal interface{}) middleware.Responder {
		return HandleStudentGet(params, principal)
	})
	api.TeamGetTeamIDHandler = team.GetTeamIDHandlerFunc(func(params team.GetTeamIDParams, principal interface{}) middleware.Responder {
		return HandleTeamGet(params, principal)
	})
	api.AddressPostAddressHandler = address.PostAddressHandlerFunc(func(params address.PostAddressParams, principal interface{}) middleware.Responder {
		return HandleAddressPost(params)
	})
	api.AdvisorPostAdvisorHandler = advisor.PostAdvisorHandlerFunc(func(params advisor.PostAdvisorParams, principal interface{}) middleware.Responder {
		return HandleAdvisorPost(params)
	})
	api.CredentialPostCredentialHandler = credential.PostCredentialHandlerFunc(func(params credential.PostCredentialParams, principal interface{}) middleware.Responder {
		return HandleCredentialPost(params)
	})
	api.ProblemPostProblemHandler = problem.PostProblemHandlerFunc(func(params problem.PostProblemParams, principal interface{}) middleware.Responder {
		return HandleProblemPost(params)
	})
	api.SchoolPostSchoolHandler = school.PostSchoolHandlerFunc(func(params school.PostSchoolParams, principal interface{}) middleware.Responder {
		return HandleSchoolPost(params)
	})
	api.SolutionPostSolutionHandler = solution.PostSolutionHandlerFunc(func(params solution.PostSolutionParams, principal interface{}) middleware.Responder {
		return HandleSolutionPost(params)
	})
	api.StudentPostStudentHandler = student.PostStudentHandlerFunc(func(params student.PostStudentParams, principal interface{}) middleware.Responder {
		return HandleStudentPost(params, principal)
	})
	api.TeamPostTeamHandler = team.PostTeamHandlerFunc(func(params team.PostTeamParams, principal interface{}) middleware.Responder {
		return HandleTeamPost(params, principal)
	})
	api.AddressPutAddressHandler = address.PutAddressHandlerFunc(func(params address.PutAddressParams, principal interface{}) middleware.Responder {
		return HandleAddressPut(params)
	})
	api.AdvisorPutAdvisorHandler = advisor.PutAdvisorHandlerFunc(func(params advisor.PutAdvisorParams, principal interface{}) middleware.Responder {
		return HandleAdvisorPut(params)
	})
	api.CredentialPutCredentialHandler = credential.PutCredentialHandlerFunc(func(params credential.PutCredentialParams, principal interface{}) middleware.Responder {
		return HandleCredentialPut(params)
	})
	api.ProblemPutProblemHandler = problem.PutProblemHandlerFunc(func(params problem.PutProblemParams, principal interface{}) middleware.Responder {
		return HandleProblemPut(params)
	})
	api.SchoolPutSchoolHandler = school.PutSchoolHandlerFunc(func(params school.PutSchoolParams, principal interface{}) middleware.Responder {
		return HandleSchoolPut(params)
	})
	api.SolutionPutSolutionHandler = solution.PutSolutionHandlerFunc(func(params solution.PutSolutionParams, principal interface{}) middleware.Responder {
		return HandleSolutionPut(params)
	})
	api.StudentPutStudentHandler = student.PutStudentHandlerFunc(func(params student.PutStudentParams, principal interface{}) middleware.Responder {
		return HandleStudentPut(params, principal)
	})
	api.TeamPutTeamHandler = team.PutTeamHandlerFunc(func(params team.PutTeamParams, principal interface{}) middleware.Responder {
		return HandleTeamPut(params, principal)
	})
	// List of all Schools
	api.SchoolGetSchoolHandler = school.GetSchoolHandlerFunc(func(params school.GetSchoolParams, principal interface{}) middleware.Responder {
		return HandleSchoolGetList(params)
	})

	// List of all Problems
	api.ProblemGetProblemHandler = problem.GetProblemHandlerFunc(func(params problem.GetProblemParams, principal interface{}) middleware.Responder {
		return HandleProblemGetList(params)
	})
	// List of all Students for a School
	api.SchoolGetSchoolIDStudentsHandler = school.GetSchoolIDStudentsHandlerFunc(func(params school.GetSchoolIDStudentsParams, principal interface{}) middleware.Responder {
		return HandleSchoolGetStudentList(params, principal)
	})
	// List all Advisors
	api.AdvisorGetAdvisorHandler = advisor.GetAdvisorHandlerFunc(func(params advisor.GetAdvisorParams, principal interface{}) middleware.Responder {
		return HandleAdvisorReadAll(params)
	})
}
