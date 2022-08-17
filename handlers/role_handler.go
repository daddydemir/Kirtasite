package handlers

import (
	"Kirtasite/services"
	"encoding/json"
	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetRoles()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
