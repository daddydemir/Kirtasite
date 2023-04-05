package handlers

import (
	"Kirtasite/auth"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var content map[string]string
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.RemoteAddr)
	code, message := auth.LoginService(content["Mail"], content["Password"], r)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
