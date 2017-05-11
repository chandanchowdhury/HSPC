package requesthandler

import (
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/restapi/operations/address"
	"github.com/go-openapi/runtime/middleware"
)

func HandleAddressPost(params address.PostAddressParams) middleware.Responder {
	//create the address
	address_id := dbhandler.AddressCreate(*params.Address)

	if address_id <= 0 {
		resp := address.NewPostAddressDefault(400)
		error := new(models.Error)
		error.Code = address_id
		error.Message = "Failed to create Address"

		if address_id == 0 {
			error.Message = "Address already exists"
		}

		resp.SetPayload(error)

		return resp
	}

	// create the response
	resp := address.NewPostAddressOK()
	address := dbhandler.AddressRead(address_id)

	//set response data
	resp.SetPayload(&address)

	//return the response
	return resp
}

func HandleAddressGet(params address.GetAddressIDParams) middleware.Responder {
	//get address details based on the provided id
	address_data := dbhandler.AddressRead(params.ID)

	if address_data.AddressID == 0 {
		resp := address.NewGetAddressIDDefault(404)
		error := &models.Error{Code: -1, Message: "Address not found"}

		resp.SetPayload(error)
		return resp
	}

	resp := address.NewGetAddressIDOK()
	resp.SetPayload(&address_data)

	return resp
}

func HandleAddressPut(params address.PutAddressParams) middleware.Responder {
	affected_count := dbhandler.AddressUpdate(*params.Address)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of updates"
		error.Code = affected_count
		resp := address.NewPostAddressDefault(400)

		//no rows updated
		if affected_count == 0 {
			error.Message = "Warn: no records found for update"
			resp.SetStatusCode(404)
		}

		resp.SetPayload(error)

		return resp
	}

	resp := address.NewPutAddressOK()
	error.Message = "Updated"
	resp.SetPayload(error)

	return resp
}

func HandleAddressDelete(params address.DeleteAddressIDParams) middleware.Responder {
	affected_count := dbhandler.AddressDelete(params.ID)

	error := new(models.Error)

	if affected_count != 1 {
		error.Message = "Error: Unexpected number of deletes"
		error.Code = affected_count
		resp := address.NewDeleteAddressIDDefault(400)

		//no rows deleted
		if affected_count == 0 {
			error.Message = "Warn: no records found for delete"
			resp.SetStatusCode(404)
		}

		resp.SetPayload(error)

		return resp
	}

	resp := address.NewDeleteAddressIDOK()
	error.Code = 0
	error.Message = "Deleted"
	resp.SetPayload(error)

	return resp
}
