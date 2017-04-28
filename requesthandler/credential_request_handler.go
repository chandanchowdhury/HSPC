package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/restapi/operations/credential"
	"github.com/go-openapi/runtime/middleware"
	"github.com/chandanchowdhury/HSPC/models"
)

func HandleCredentialPost(params credential.PostCredentialParams) middleware.Responder {
	//create the credential
	credential_id := dbhandler.CredentialCreate(*params.Body)

	// create the response
	resp := credential.NewPostCredentialOK()
	error := new(models.Error)

	error.Code = credential_id
	error.Message = "Created"

	//set response data
	resp.SetPayload(error)

	//return the response
	return resp
}

func HandleCredentialGet(params credential.GetCredentialIDParams) middleware.Responder {
	//get credential details based on the provided id
	credential_data := dbhandler.CredentialRead(params.ID.String())

	//credential not found
	if credential_data.CredentialID == 0 {
		resp := credential.NewGetCredentialIDDefault(404)
		error := &models.Error{Code: -1, Message:"User not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := credential.NewGetCredentialIDOK()

	resp.SetPayload(credential_data)

	return resp
}

func HandleCredentialPut(params credential.PutCredentialParams) middleware.Responder {
	affected_count := dbhandler.CredentialUpdate(params.Body.Emailaddress.String(), params.Body.Password.String(), *params.Body.CredentialActive)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := credential.NewPutCredentialDefault(-1)
		resp.SetPayload(error)

		return resp
	}

	error.Message = "Updated"
	resp := credential.NewPutCredentialOK()
	resp.SetPayload(error)

	return resp
}

func HandleCredentialDelete(params credential.DeleteCredentialIDParams) middleware.Responder {
	affected_count := dbhandler.CredentialDelete(params.ID.String())

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := credential.NewDeleteCredentialIDDefault(-1)
		resp.SetPayload(error)

		return resp
	}

	error.Message = "Deleted"
	resp := credential.NewDeleteCredentialIDOK()
	resp.SetPayload(error)

	return resp
}
