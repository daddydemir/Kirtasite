package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetCustomers()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetCustomerByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetCustomerByUserId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var customer models.Customers
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &customer)

	code, message := services.AddCustomer(customer)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var customer models.Customers
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &customer)

	code, message := services.UpdateCustomer(customer, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
