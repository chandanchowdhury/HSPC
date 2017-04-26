package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/restapi/operations/credential"
	"github.com/go-openapi/runtime/middleware"
)

func HandleCredentialPost(params credential.PostCredentialParams) middleware.Responder {
	//create the credential
	_ = dbhandler.CredentialCreate(*params.Body)

	credential_data := dbhandler.CredentialRead(params.Body.Emailaddress.String())

	// create the response
	resp := credential.NewGetCredentialIDOK()

	//set response data
	resp.SetPayload(&credential_data)

	//return the response
	return resp
}

func HandleCredentialGet(params credential.GetCredentialIDParams) middleware.Responder {
	//get credential details based on the provided id
	credential_data := dbhandler.CredentialRead(params.ID.String())

	resp := credential.NewGetCredentialIDOK()

	resp.SetPayload(&credential_data)

	return resp
}

func HandleCredentialPut(params credential.PutCredentialParams) middleware.Responder {
	affected_count := dbhandler.CredentialUpdate(params.Body.Emailaddress.String(), params.Body.Password.String())

	var resp_str string
	resp_str = "Success"

	if affected_count != 1 {
		resp_str = "Error"
	}

	resp := credential.NewPutCredentialOK()
	resp.SetPayload(resp_str)

	return resp
}

func HandleCredentialDelete(params credential.DeleteCredentialIDParams) middleware.Responder {
	affected_count := dbhandler.CredentialDelete(params.ID.String())

	var resp_str string
	resp_str = "Success"

	if affected_count != 1 {
		resp_str = "Error"
	}

	resp := credential.NewDeleteCredentialIDOK()
	resp.SetPayload(resp_str)

	return resp
}
