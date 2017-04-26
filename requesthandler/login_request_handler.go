package requesthandler

import (
    "github.com/go-openapi/runtime/middleware"
    "github.com/chandanchowdhury/HSPC/restapi/operations/login"
    "github.com/chandanchowdhury/HSPC/dbhandler"
    "log"
)

func HandleLogin(params login.GetLoginEmailaddressPasswordParams) middleware.Responder {
    email := params.Emailaddress.String()
    password := params.Password.String()

    log.Printf("Received - email: %s, password: %s", email, password)

    credential := dbhandler.CredentialRead(email)

    //log.Printf("In DB - email: %p, password: %p", credential.Emailaddress, credential.Password)
    log.Printf("In DB - email: %s, password: %s Active: %b", credential.Emailaddress.String(), credential.Password.String(), credential.CredentialActive)

    resp := login.NewGetLoginEmailaddressPasswordOK()
    var session_id string

    session_id = "Failed"

    //TODO: if the credential is not active, login is not allowed
    //if credential.Emailaddress != nil && credential.CredentialActive == false {
    //    session_id = "Not Active"
    //    resp.SetPayload(session_id)
    //    return resp
    //}

    if credential.Emailaddress != nil && credential.Password.String() == password {
                //TODO: Setup session
                session_id = "Success"
    }

    resp.SetPayload(session_id)
    return resp
}
