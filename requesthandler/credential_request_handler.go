package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/credential"
	"github.com/go-openapi/runtime/middleware"
)

func HandleCredentialPost(params credential.PostCredentialParams) middleware.Responder {
	//create the credential
	credential_id := dbhandler.CredentialCreate(*params.Credential)

	if credential_id <= 0 {
		resp := credential.NewPostCredentialDefault(400)
		error := new(models.Error)
		error.Code = credential_id
		error.Message = "Error: Failed to create Credential"

		if credential_id == -1 {
			error.Message = "Credential already exists"
		}

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := credential.NewPostCredentialOK()
	credential := dbhandler.CredentialRead(params.Credential.Emailaddress.String())
	//set response data
	resp.SetPayload(credential)

	//return the response
	return resp
}

func HandleCredentialGet(params credential.GetCredentialEmailaddressParams) middleware.Responder {
	//get credential details based on the provided id
	credential_data := dbhandler.CredentialRead(params.Emailaddress.String())

	//credential not found
	if credential_data.CredentialID == 0 {
		resp := credential.NewGetCredentialEmailaddressDefault(404)
		error := &models.Error{Code: -1, Message: "User not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := credential.NewGetCredentialEmailaddressOK()

	resp.SetPayload(credential_data)

	return resp
}

func HandleCredentialPut(params credential.PutCredentialParams) middleware.Responder {
	affected_count := dbhandler.CredentialUpdate(params.Credential.Emailaddress.String(), params.Credential.Password.String(), *params.Credential.CredentialActive)

	error := new(models.Error)

	if affected_count <= 0 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := credential.NewPutCredentialDefault(400)

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

	error.Message = "Updated"
	resp := credential.NewPutCredentialOK()
	resp.SetPayload(error)

	return resp
}

func HandleCredentialDelete(params credential.DeleteCredentialEmailaddressParams) middleware.Responder {
	affected_count := dbhandler.CredentialDelete(params.Emailaddress.String())

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := credential.NewDeleteCredentialEmailaddressDefault(400)

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

	error.Message = "Deleted"
	resp := credential.NewDeleteCredentialEmailaddressOK()
	resp.SetPayload(error)

	return resp
}
