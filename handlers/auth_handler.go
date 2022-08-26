package handlers

import (
	"Kirtasite/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var content map[string]string
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &content)

	code, message := auth.Login(content["Mail"], content["Password"])
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
