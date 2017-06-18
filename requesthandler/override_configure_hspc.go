package requesthandler

import (
	"log"

	"time"

	"net/http"

	"bytes"
	"encoding/base64"
	"encoding/gob"

	"crypto/hmac"
	"crypto/sha256"

	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations"
	"github.com/chandanchowdhury/HSPC/restapi/operations/address"
	"github.com/chandanchowdhury/HSPC/restapi/operations/advisor"
	"github.com/chandanchowdhury/HSPC/restapi/operations/credential"
	"github.com/chandanchowdhury/HSPC/restapi/operations/login"
	"github.com/chandanchowdhury/HSPC/restapi/operations/logout"
	"github.com/chandanchowdhury/HSPC/restapi/operations/problem"
	"github.com/chandanchowdhury/HSPC/restapi/operations/school"
	"github.com/chandanchowdhury/HSPC/restapi/operations/solution"
	"github.com/chandanchowdhury/HSPC/restapi/operations/student"
	"github.com/chandanchowdhury/HSPC/restapi/operations/team"
	"github.com/didip/tollbooth"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
)

//TODO: Generate 128 bit random hash key
var hashKey = []byte("very-secret1234")

/**
This function when called inside configureAPI() function in
 configure_hspc.go will override all NotImplemented calls with actual
 valid calls to handle the request.

Just add below line before return in configureAPI()

requesthandler.Override_configure_hspc(api)
*/
func Override_configure_hspc(api *operations.HspcAPI) {

	// Applies when the "api-key" header is set
	api.APISecurityAuth = func(token string) (*models.Principal, error) {
		log.Printf("Checking API key = %s", token)

		//TODO: need better API key validation
		if token == "This is ADMIN app API key" {
			prin := new(models.Principal)
			prin.Email = "Admin"
			log.Print("Request using ADMIN API key")
			return prin, nil
		}

		if token == "This is USER app API key" {
			prin := new(models.Principal)
			prin.Email = "User"

			log.Print("Request using USER API key")
			return prin, nil
		}

		log.Print("No or Invalid API key")

		return nil, errors.New(403, "Invalid API key")
	}

	// Applies when the "Hspc-Token" header is set
	api.SessionSecurityAuth = func(token string) (*models.Principal, error) {
		log.Printf("SessionSecurityAuth - Token = %s", token)

		//convert the binary gob to base64 URL encoding
		uDec, err := base64.URLEncoding.DecodeString(token)
		if err != nil {
			log.Print("Base64 decoding failed")
			return nil, errors.New(401, "Session Not Found")
		}

		// deserialize the GOB
		// create a buffer
		in := new(bytes.Buffer)
		//write the token data into the buffer as bytes
		in.Write([]byte(uDec))
		// create a decoder to decode the data
		e := gob.NewDecoder(in)

		prin := new(models.Principal)
		err = e.Decode(&prin)
		if err != nil {
			log.Print("GOB decoding failed")
			log.Panic(err)
		}

		//Check that the HMAC is valid. To check recalculate the HMAC of email and Timestamp and compare
		// with the value sent by the client.
		hmac_func := hmac.New(sha256.New, hashKey)

		//add email
		hmac_func.Write([]byte(prin.Email))
		//add Timestamp
		hmac_func.Write([]byte(prin.CreatedTs))
		// encode the HMAC in Base64 and store it
		sessionToken := base64.URLEncoding.EncodeToString(hmac_func.Sum(nil))

		//does the session exists in our database?
		prin_in_db := dbhandler.SessionRead(prin.Email)

		//compare the HMAC generated against the HMAC sent by the client
		if sessionToken == prin.SessionToken && prin_in_db.Email == prin.Email {
			log.Printf("Session Principal: %s", prin.Email)
			return prin, nil
		}

		return nil, errors.NotFound("Session not found")
	}

	api.LoginPostLoginHandler = login.PostLoginHandlerFunc(func(params login.PostLoginParams, principal *models.Principal) middleware.Responder {
		return HandleLogin(params, principal)
	})

	api.LogoutPostLogoutHandler = logout.PostLogoutHandlerFunc(func(params logout.PostLogoutParams, principal *models.Principal) middleware.Responder {
		return HandleLogout(params, principal)
	})

	api.AddressDeleteAddressIDHandler = address.DeleteAddressIDHandlerFunc(func(params address.DeleteAddressIDParams, principal *models.Principal) middleware.Responder {
		return HandleAddressDelete(params, principal)
	})
	api.AdvisorDeleteAdvisorIDHandler = advisor.DeleteAdvisorIDHandlerFunc(func(params advisor.DeleteAdvisorIDParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorDelete(params)
	})
	api.CredentialDeleteCredentialEmailaddressHandler = credential.DeleteCredentialEmailaddressHandlerFunc(func(params credential.DeleteCredentialEmailaddressParams, principal *models.Principal) middleware.Responder {
		return HandleCredentialDelete(params)
	})
	api.ProblemDeleteProblemIDHandler = problem.DeleteProblemIDHandlerFunc(func(params problem.DeleteProblemIDParams, principal *models.Principal) middleware.Responder {
		return HandleProblemDelete(params)
	})
	api.SchoolDeleteSchoolIDHandler = school.DeleteSchoolIDHandlerFunc(func(params school.DeleteSchoolIDParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolDelete(params, principal)
	})
	api.SolutionDeleteSolutionIDHandler = solution.DeleteSolutionIDHandlerFunc(func(params solution.DeleteSolutionIDParams, principal *models.Principal) middleware.Responder {
		return HandleSolutionDelete(params)
	})
	api.StudentDeleteStudentIDHandler = student.DeleteStudentIDHandlerFunc(func(params student.DeleteStudentIDParams, principal *models.Principal) middleware.Responder {
		return HandleStudentDelete(params, principal)
	})
	api.TeamDeleteTeamIDHandler = team.DeleteTeamIDHandlerFunc(func(params team.DeleteTeamIDParams, principal *models.Principal) middleware.Responder {
		return HandleTeamDelete(params, principal)
	})
	api.AddressGetAddressIDHandler = address.GetAddressIDHandlerFunc(func(params address.GetAddressIDParams, principal *models.Principal) middleware.Responder {
		return HandleAddressGet(params)
	})
	api.AdvisorGetAdvisorIDHandler = advisor.GetAdvisorIDHandlerFunc(func(params advisor.GetAdvisorIDParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorGet(params)
	})
	api.CredentialGetCredentialEmailaddressHandler = credential.GetCredentialEmailaddressHandlerFunc(func(params credential.GetCredentialEmailaddressParams, principal *models.Principal) middleware.Responder {
		return HandleCredentialGet(params)
	})

	api.ProblemGetProblemIDHandler = problem.GetProblemIDHandlerFunc(func(params problem.GetProblemIDParams, principal *models.Principal) middleware.Responder {
		return HandleProblemGet(params)
	})
	api.SchoolGetSchoolIDHandler = school.GetSchoolIDHandlerFunc(func(params school.GetSchoolIDParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolGet(params, principal)
	})
	api.SolutionGetSolutionIDHandler = solution.GetSolutionIDHandlerFunc(func(params solution.GetSolutionIDParams, principal *models.Principal) middleware.Responder {
		return HandleSolutionGet(params)
	})
	api.StudentGetStudentIDHandler = student.GetStudentIDHandlerFunc(func(params student.GetStudentIDParams, principal *models.Principal) middleware.Responder {
		return HandleStudentGet(params, principal)
	})
	api.TeamGetTeamIDHandler = team.GetTeamIDHandlerFunc(func(params team.GetTeamIDParams, principal *models.Principal) middleware.Responder {
		return HandleTeamGet(params, principal)
	})
	api.AddressPostAddressHandler = address.PostAddressHandlerFunc(func(params address.PostAddressParams, principal *models.Principal) middleware.Responder {
		return HandleAddressPost(params)
	})
	api.AdvisorPostAdvisorHandler = advisor.PostAdvisorHandlerFunc(func(params advisor.PostAdvisorParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorPost(params)
	})
	api.CredentialPostCredentialHandler = credential.PostCredentialHandlerFunc(func(params credential.PostCredentialParams, principal *models.Principal) middleware.Responder {
		return HandleCredentialPost(params)
	})
	api.ProblemPostProblemHandler = problem.PostProblemHandlerFunc(func(params problem.PostProblemParams, principal *models.Principal) middleware.Responder {
		return HandleProblemPost(params)
	})
	api.SchoolPostSchoolHandler = school.PostSchoolHandlerFunc(func(params school.PostSchoolParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolPost(params, principal)
	})
	api.SolutionPostSolutionHandler = solution.PostSolutionHandlerFunc(func(params solution.PostSolutionParams, principal *models.Principal) middleware.Responder {
		return HandleSolutionPost(params)
	})
	api.StudentPostStudentHandler = student.PostStudentHandlerFunc(func(params student.PostStudentParams, principal *models.Principal) middleware.Responder {
		return HandleStudentPost(params, principal)
	})
	api.TeamPostTeamHandler = team.PostTeamHandlerFunc(func(params team.PostTeamParams, principal *models.Principal) middleware.Responder {
		return HandleTeamPost(params, principal)
	})
	api.AddressPutAddressHandler = address.PutAddressHandlerFunc(func(params address.PutAddressParams, principal *models.Principal) middleware.Responder {
		return HandleAddressPut(params)
	})
	api.AdvisorPutAdvisorHandler = advisor.PutAdvisorHandlerFunc(func(params advisor.PutAdvisorParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorPut(params)
	})
	api.CredentialPutCredentialHandler = credential.PutCredentialHandlerFunc(func(params credential.PutCredentialParams, principal *models.Principal) middleware.Responder {
		return HandleCredentialPut(params)
	})
	api.ProblemPutProblemHandler = problem.PutProblemHandlerFunc(func(params problem.PutProblemParams, principal *models.Principal) middleware.Responder {
		return HandleProblemPut(params)
	})
	api.SchoolPutSchoolHandler = school.PutSchoolHandlerFunc(func(params school.PutSchoolParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolPut(params, principal)
	})
	api.SolutionPutSolutionHandler = solution.PutSolutionHandlerFunc(func(params solution.PutSolutionParams, principal *models.Principal) middleware.Responder {
		return HandleSolutionPut(params)
	})
	api.StudentPutStudentHandler = student.PutStudentHandlerFunc(func(params student.PutStudentParams, principal *models.Principal) middleware.Responder {
		return HandleStudentPut(params, principal)
	})
	api.TeamPutTeamHandler = team.PutTeamHandlerFunc(func(params team.PutTeamParams, principal *models.Principal) middleware.Responder {
		return HandleTeamPut(params, principal)
	})

	// --- School ---
	// List of all Schools
	api.SchoolGetSchoolHandler = school.GetSchoolHandlerFunc(func(params school.GetSchoolParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolGetList(params, principal)
	})
	// List of all Students for a School
	api.SchoolGetSchoolIDStudentsHandler = school.GetSchoolIDStudentsHandlerFunc(func(params school.GetSchoolIDStudentsParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolGetStudentList(params, principal)
	})
	//List all Teams for a School
	api.SchoolGetSchoolIDTeamsHandler = school.GetSchoolIDTeamsHandlerFunc(func(params school.GetSchoolIDTeamsParams, principal *models.Principal) middleware.Responder {
		return HandleSchoolGetTeamList(params, principal)
	})
	// Add School Advisor
	api.SchoolPutSchoolSchoolIDAdvisorAdvisorIDHandler = school.PutSchoolSchoolIDAdvisorAdvisorIDHandlerFunc(func(params school.PutSchoolSchoolIDAdvisorAdvisorIDParams, principal *models.Principal) middleware.Responder {
		return SchoolAddAdvisor(params, principal)
	})

	// Remove School Advisor
	api.SchoolDeleteSchoolSchoolIDAdvisorAdvisorIDHandler = school.DeleteSchoolSchoolIDAdvisorAdvisorIDHandlerFunc(func(params school.DeleteSchoolSchoolIDAdvisorAdvisorIDParams, principal *models.Principal) middleware.Responder {
		return SchoolRemoveAdvisor(params, principal)
	})

	// --- Advisor ---
	// List all Advisors
	api.AdvisorGetAdvisorHandler = advisor.GetAdvisorHandlerFunc(func(params advisor.GetAdvisorParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorReadAll(params)
	})
	// List all Schools for Advisor
	api.AdvisorGetAdvisorIDSchoolsHandler = advisor.GetAdvisorIDSchoolsHandlerFunc(func(params advisor.GetAdvisorIDSchoolsParams, principal *models.Principal) middleware.Responder {
		return HandleAdvisorListSchool(params, principal)
	})

	// List of all Problems
	api.ProblemGetProblemHandler = problem.GetProblemHandlerFunc(func(params problem.GetProblemParams, principal *models.Principal) middleware.Responder {
		return HandleProblemGetList(params)
	})
	//List all solutions for a problem
	api.ProblemGetProblemIDSolutionsHandler = problem.GetProblemIDSolutionsHandlerFunc(func(params problem.GetProblemIDSolutionsParams, principal *models.Principal) middleware.Responder {
		return HandleGetProblemIDSolutions(params)
	})

	// --- Team ---
	// Get Team members
	api.TeamGetTeamIDStudentsHandler = team.GetTeamIDStudentsHandlerFunc(func(params team.GetTeamIDStudentsParams, principal *models.Principal) middleware.Responder {
		return HandleGetTeamStudents(params, principal)
	})
	//Add Team member
	api.TeamGetTeamTeamIDAddStudentIDHandler = team.GetTeamTeamIDAddStudentIDHandlerFunc(func(params team.GetTeamTeamIDAddStudentIDParams, principal *models.Principal) middleware.Responder {
		return HandleTeamAddStudent(params, principal)
	})
	//Remove Team member
	api.TeamGetTeamTeamIDRemoveStudentIDHandler = team.GetTeamTeamIDRemoveStudentIDHandlerFunc(func(params team.GetTeamTeamIDRemoveStudentIDParams, principal *models.Principal) middleware.Responder {
		return HandleTeamRemoveStudent(params, principal)
	})
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	log.Print("RateLimitMiddleware()")
	limiter := tollbooth.NewLimiter(5, time.Second)
	limiter.IPLookups = []string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}
	return tollbooth.LimitFuncHandler(limiter, next.ServeHTTP)
}

func HandleCORS(next http.Handler) http.Handler {
	log.Print("HandleCORS()")

	//set up the CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Hspc-Api-Key", "Hspc-Token", "Emailaddress", "Password"},
		ExposedHeaders: []string{"Hspc-Token"},
		//Debug:          true,
	})

	return c.Handler(RateLimitMiddleware(next))
	//return c.Handler(next)
}

/**
Call this function inside setupGlobalMiddleware in configure_hspc.go to setup CORS and RateLimit handling middlewares

return requesthandler.HspcMiddlewares(handler)
*/
func HspcMiddlewares(next http.Handler) http.Handler {
	return HandleCORS(RateLimitMiddleware(next))
}
