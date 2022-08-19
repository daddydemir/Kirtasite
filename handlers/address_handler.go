package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetAddressById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetAddressById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetAddressByCityId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetAddressByCityId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func AddAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var address models.Addresses
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &address)

	code, message := services.AddAddress(address)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var address models.Addresses
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &address)

	code, message := services.UpdateAddress(address, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
