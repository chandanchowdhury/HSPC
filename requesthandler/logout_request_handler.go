package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/logout"
	"github.com/go-openapi/runtime/middleware"
)

func HandleLogout(params logout.PostLogoutParams, principal *models.Principal) middleware.Responder {
	//if valid principal, return logout okay
	if principal != nil {
		resp := logout.NewPostLogoutOK()
		return resp
	}

	//invalid principal, logout unsuccessful
	resp := logout.NewPostLogoutDefault(401)
	return resp
}
