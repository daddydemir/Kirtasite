package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetStationeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetStationeries()
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func GetStationeryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetStationeryById(key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func AddStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var stationery models.Stationeries
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &stationery)

	code, message := services.AddStationery(stationery)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateStationery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var stationery models.Stationeries
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &stationery)

	code, message := services.UpdateStationery(stationery, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
