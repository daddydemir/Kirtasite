package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	code, message := services.GetPrices()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetPriceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetPriceById(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func GetPriceByStationeryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetPriceByStationeryId(key)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func AddPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var price models.Price
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &price)

	code, message := services.AddPrice(price)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
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

	var price models.Price
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &price)

	code, message := services.UpdatePrice(price, key, t)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
