package handlers

import (
	"Kirtasite/models"
	"Kirtasite/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetOrderById(key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func GetOrdersByCustomerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetOrdersByCustomerId(key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func GetOrdersByStationeryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, GET)

	vars := mux.Vars(r)
	key := vars["id"]

	code, message := services.GetOrdersByStationeryId(key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, POST)

	var order models.Orders
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &order)

	code, message := services.AddOrder(order)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JSON)
	w.Header().Set(AccessOrigin, ORIGIN)
	w.Header().Set(AccessMethods, PUT)

	vars := mux.Vars(r)
	key := vars["id"]

	var order models.Orders
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &order)

	code, message := services.UpdateOrder(order, key)
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
