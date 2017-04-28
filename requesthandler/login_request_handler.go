package requesthandler

import (
    "github.com/go-openapi/runtime/middleware"
    "github.com/chandanchowdhury/HSPC/restapi/operations/login"
    "github.com/chandanchowdhury/HSPC/dbhandler"
    "github.com/chandanchowdhury/HSPC/models"
    "log"
)

func HandleLogin(params login.GetLoginEmailaddressPasswordParams) middleware.Responder {
    email := params.Emailaddress.String()
    password := params.Password.String()
    log.Printf("Received - email: %s, password: %s", email, password)

    error := new(models.Error)
    error.Message = "Failed"

    credential := dbhandler.CredentialRead(email)

    if credential.Emailaddress == nil {
        resp := login.NewGetLoginEmailaddressPasswordOK()
        resp.SetPayload(error)
        return resp
    }

    log.Printf("From DB - email: %s, password: %s Active: %t", credential.Emailaddress.String(), credential.Password.String(), *credential.CredentialActive)

    if credential.Password.String() == password {
        log.Print("Password Matched")
        error.Message = "Success"
    }

    if *credential.CredentialActive == false {
        log.Print("Account Not Active")
        error.Message = "Acount Not Active"
    }

    resp := login.NewGetLoginEmailaddressPasswordOK()
    resp.SetPayload(error)
    return resp
}
