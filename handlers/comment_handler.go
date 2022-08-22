package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetCommentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetCommentById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetCommentsByStationeryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetCommentsByStationeryId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetCommentsByCustomerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetCommentsByCustomerId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var comment models.Comments
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &comment)

	code, message := services.AddComment(comment)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var comment models.Comments
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &comment)

	code, message := services.UpdateComment(comment, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
