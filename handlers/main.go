package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func MainRouting() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	baseurl := "/api/v1/"

	// Roles
	r.HandleFunc(baseurl+"roles", GetRoles).Methods(GET)

	// Cities
	r.HandleFunc(baseurl+"cities", GetCities).Methods(GET)
	r.HandleFunc(baseurl+"city/{id:[0-9]+}", GetCityById).Methods(GET)

	handler := cors.AllowAll().Handler(r)
	return handler
}
