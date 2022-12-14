package handlers

import (
	"Kirtasite/cloudinary"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	s, m, t := _tokenCheck(r)
	if !s {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	// get file
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}
	url, err := cloudinary.Upload(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}
	// image uploaded

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.ImageUpload(url, key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func PasswordUpdate(w http.ResponseWriter, r *http.Request) {
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

	var content map[string]string
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &content)

	code, message := services.PasswordUpdate(content["Lastpass"], content["Againpass"], content["Newpass"], key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
