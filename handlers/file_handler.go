package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetFilesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetFilesById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetFilesByCustomerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetFilesByCustomerId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func AddFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var file models.Files
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &file)

	code, message := services.AddFiles(file)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var file models.Files
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &file)

	code, message := services.UpdateFiles(file, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}