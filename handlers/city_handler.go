package handlers

import (
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetCities()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetCityById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetCityById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
