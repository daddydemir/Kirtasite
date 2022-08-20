package handlers

import (
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetOrderStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetOrderStatuses()
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func GetOrderStatusById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetOrderStatusById(key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
