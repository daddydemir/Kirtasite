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

	// Customers
	r.HandleFunc(baseurl+"customers", GetCustomers).Methods(GET)
	r.HandleFunc(baseurl+"customer/{id:[0-9]+}", GetCustomerByUserId).Methods(GET)
	r.HandleFunc(baseurl+"customer", AddCustomer).Methods(POST)
	r.HandleFunc(baseurl+"customer/{id:[0-9]+}", UpdateCustomer).Methods(PUT)

	// Stationeries
	r.HandleFunc(baseurl+"stationeries", GetStationeries).Methods(GET)
	r.HandleFunc(baseurl+"stationery/{id:[0-9]+}", GetStationeryById).Methods(GET)
	r.HandleFunc(baseurl+"stationery", AddStationery).Methods(POST)
	r.HandleFunc(baseurl+"stationery/{id:[0-9]+}", UpdateStationery).Methods(PUT)

	// Prices
	r.HandleFunc(baseurl+"prices", GetPrices).Methods(GET)
	r.HandleFunc(baseurl+"price/{id:[0-9]+}", GetPriceById).Methods(GET)
	r.HandleFunc(baseurl+"prices/stationery/{id:[0-9]+}", GetPriceByStationeryId).Methods(GET)
	r.HandleFunc(baseurl+"price", AddPrice).Methods(POST)
	r.HandleFunc(baseurl+"price/{id:[0-9]+}", UpdatePrice).Methods(PUT)

	// Orders
	r.HandleFunc(baseurl+"order/{id:[0-9]+}", GetOrderById).Methods(GET)
	r.HandleFunc(baseurl+"orders/customer/{id:[0-9]+}", GetOrdersByCustomerId).Methods(GET)
	r.HandleFunc(baseurl+"orders/stationery/{id:[0-9]+}", GetOrdersByStationeryId).Methods(GET)
	r.HandleFunc(baseurl+"order", AddOrder).Methods(POST)
	// Artık kullanılmıyor
	//r.HandleFunc(baseurl+"order/{id:[0-9]+}", UpdateOrder).Methods(PUT)
	r.HandleFunc(baseurl+"order/{id:[0-9]+}/cancel", CancelOrder).Methods(PUT)
	r.HandleFunc(baseurl+"order/{id:[0-9]+}/confirm", ConfirmOrder).Methods(PUT)
	r.HandleFunc(baseurl+"order/{id:[0-9]+}/ready", ReadyOrder).Methods(PUT)
	r.HandleFunc(baseurl+"order/{id:[0-9]+}/complete", CompleteOrder).Methods(PUT)

	// OrderStatuses
	r.HandleFunc(baseurl+"statuses", GetOrderStatuses).Methods(GET)
	r.HandleFunc(baseurl+"status/{id:[0-9]+}", GetOrderStatusById).Methods(GET)

	// Files
	r.HandleFunc(baseurl+"file/{id:[0-9]+}", GetFilesById).Methods(GET)
	r.HandleFunc(baseurl+"files/customer/{id:[0-9]+}", GetFilesByCustomerId).Methods(GET)
	r.HandleFunc(baseurl+"file", AddFiles).Methods(POST)
	r.HandleFunc(baseurl+"file/{id:[0-9]+}", UpdateFiles).Methods(PUT)

	// Comments
	r.HandleFunc(baseurl+"comment/{id:[0-9]+}", GetCommentById).Methods(GET)
	r.HandleFunc(baseurl+"comments/stationery/{id:[0-9]+}", GetCommentsByStationeryId).Methods(GET)
	r.HandleFunc(baseurl+"comments/customer/{id:[0-9]+}", GetCommentsByCustomerId).Methods(GET)
	r.HandleFunc(baseurl+"comment", AddComment).Methods(POST)
	r.HandleFunc(baseurl+"comment/{id:[0-9]+}", UpdateComment).Methods(PUT)

	// Login
	r.HandleFunc(baseurl+"login", Login).Methods(POST)

	// Users
	r.HandleFunc(baseurl+"user/image/{id:[0-9]+}", ImageUpload).Methods(PUT)
	r.HandleFunc(baseurl+"user/password/{id:[0-9]+}", PasswordUpdate).Methods(PUT)

	handler := cors.AllowAll().Handler(r)
	return handler
}
