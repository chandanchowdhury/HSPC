package requesthandler

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"log"
	"time"

	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/login"
	"github.com/go-openapi/runtime/middleware"
)

func HandleLogin(params login.PostLoginParams, principal *models.Principal) middleware.Responder {
	log.Print("api.LoginPostLoginHandler()")

	user := params.Emailaddress.String()
	pass := params.Password.String()

	log.Printf("Authenticating User = %s", user)

	credential := dbhandler.CredentialRead(user)

	if credential.CredentialID == 0 {
		resp := login.NewPostLoginDefault(401)
		return resp
	}

	if credential.Password.String() == pass && *credential.CredentialActive == true {
		log.Print("Password Matched")
		resp := login.NewPostLoginOK()

		prin := new(models.Principal)
		prin.Email = user
		ts := time.Now()
		prin.CreatedTs = ts.Format(time.UnixDate)

		hmac_func := hmac.New(sha256.New, hashKey)

		//add email
		hmac_func.Write([]byte(prin.Email))
		//add Timestamp
		hmac_func.Write([]byte(prin.CreatedTs))
		// encode the HMAC in Base64 and save it in SessionToken
		prin.SessionToken = base64.URLEncoding.EncodeToString(hmac_func.Sum(nil))
		// alternate to Base64 ecnoding
		//prin.SessionToken = fmt.Sprintf("%x", hmac_func.Sum(nil))

		//Save the session data in a database to improve security
		serr := dbhandler.SessionCreate(prin)

		if serr.Code != 0 {
			log.Print("Failed to save session in database")
			log.Panic(serr)
		}

		//now serialize the Principal object and return to API caller
		// gob encoder
		b := new(bytes.Buffer)
		e := gob.NewEncoder(b)
		err := e.Encode(prin)

		if err != nil {
			log.Panic("GOB Encoding failed")
		}

		//convert the binary gob to base64 URL encoding
		uEnc := base64.URLEncoding.EncodeToString([]byte(b.Bytes()))

		//log.Printf("SessionID = %s", uEnc)

		resp.HspcToken = uEnc
		return resp
	}

	resp := login.NewPostLoginDefault(401)
	return resp
}
