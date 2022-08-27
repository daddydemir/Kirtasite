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

	var comment models.Comment
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

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	var comment models.Comment
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &comment)

	code, message := services.UpdateComment(comment, key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
