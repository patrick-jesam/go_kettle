package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a accountAPI) mw(ct http.HandlerFunc) {

}

func (a accountAPI) Router(router *mux.Router) {
	userRoute := router.PathPrefix("/v1/plan").Subrouter()

	// get
	userRoute.HandleFunc("", a.GetAccount).Methods("GET")
}
