package response

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	// C serves as a placeholder for the http request and response
	C struct {
		W http.ResponseWriter
		R *http.Request
	}

	// GeneralResponse returns a standard response format
	GeneralResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
		Token   string      `json:"token,omitempty"`
		Error   interface{} `json:"error,omitempty"`
	}

	// H defines a json type format
	H map[string]interface{}
)

func (c *C) Params(key string) string {
	return mux.Vars(c.R)[key]
}

func (c *C) BindJSON(data interface{}) {
	err := json.NewDecoder(c.R.Body).Decode(data)

	if err != nil {
		return
	}
}

// Response returns a json response to the requester
func (c *C) Response200(data interface{}, message string) {
	responseSuccess := GeneralResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
	c.responseJSON(c.W, http.StatusOK, responseSuccess)
}

func (c *C) Response401(res http.ResponseWriter, resp GeneralResponse) {
	c.responseJSON(res, http.StatusUnauthorized, H{"error": resp.Error, "status": false, "message": resp.Message})
}

func (c *C) Response200WithToken(data interface{}, token, message string) {
	responseSuccess := GeneralResponse{
		Success: true,
		Data:    data,
		Message: message,
		Token:   token,
	}
	c.responseJSON(c.W, http.StatusOK, responseSuccess)
}

func (c *C) Response(success bool, data interface{}, message string, status int) {
	responseSuccess := GeneralResponse{
		Success: success,
		Data:    data,
		Message: message,
	}
	c.responseJSON(c.W, status, responseSuccess)
}

func (c *C) responseJSON(res http.ResponseWriter, status int, object interface{}) {
	res.Header().Set("Content-Resource", "application/json")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	if err := json.NewEncoder(res).Encode(object); err != nil {
		return
	}
}

func (c *C) ResponseError(error error, action string) {
	response := GeneralResponse{
		Success: false,
		Message: error.Error(),
	}

	response.Message = "unable to process request at this time"

	log.Println(action, error)

	c.responseJSON(c.W, http.StatusInternalServerError, response)
	return
}

func (c *C) ResponseValidateErr(error error) {
	response := GeneralResponse{
		Success: false,
		Message: error.Error(),
	}

	c.responseJSON(c.W, http.StatusBadRequest, response)
	return
}
