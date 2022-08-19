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

	// Districts
	r.HandleFunc(baseurl+"districts", GetDistricts).Methods(GET)
	r.HandleFunc(baseurl+"district/city/{id:[0-9]+}", GetDistrictByCityId).Methods(GET)
	r.HandleFunc(baseurl+"district/{id:[0-9]+}", GetDistrictById).Methods(GET)

	// Addresses
	r.HandleFunc(baseurl+"address/{id:[0-9]+}", GetAddressById).Methods(GET)
	r.HandleFunc(baseurl+"addresses/city/{id:[0-9]+}", GetAddressByCityId).Methods(GET)
	r.HandleFunc(baseurl+"address", AddAddress).Methods(POST)
	r.HandleFunc(baseurl+"address/{id:[0-9]+}", UpdateAddress).Methods(PUT)

	handler := cors.AllowAll().Handler(r)
	return handler
}
