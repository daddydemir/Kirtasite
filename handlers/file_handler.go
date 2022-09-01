package handlers

import (
	"Kirtasite/cloud"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetFilesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetFilesById(key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func GetFilesByCustomerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]
	code, message := services.GetFilesByCustomerId(key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func AddFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	file := models.File{}
	file.Id, _ = strconv.Atoi(r.FormValue("id"))
	file.CustomerId, _ = strconv.Atoi(r.FormValue("customer_id"))
	if r.FormValue("private") == "true" {
		file.Private = true
	} else {
		file.Private = false
	}

	content := r.ContentLength
	if content >= 1024*1024*5 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.SendMessage(core.FileSize))
		return
	}
	myfile, header, _ := r.FormFile("file")
	defer myfile.Close()
	path, folder, err := cloud.UploadFile(myfile, header)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(core.SendMessage(core.ServiceError))
		return
	}
	file.FilePath = path
	file.FolderId = folder
	file.CreatedDate = time.Now()

	code, message := services.AddFiles(file, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func UpdateFiles(w http.ResponseWriter, r *http.Request) {
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

	var file models.File
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &file)

	code, message := services.UpdateFiles(file, key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
