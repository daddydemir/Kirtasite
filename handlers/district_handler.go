package handlers

import (
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetDistricts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetDistricts()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetDistrictByCityId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetDistrictByCityId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetDistrictById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetDistrictById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
