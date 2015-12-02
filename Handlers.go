package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	allRoutes := GetAllRows("routes")

	if err := json.NewEncoder(w).Encode(allRoutes); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromAirport := strings.Trim(vars["fromAirport"], " ")
	toAirport := strings.Trim(vars["toAirport"], " ")
	routeInfo := GetRowsMulti("routes", "SourceAirportCode", fromAirport, "DestAirportCode", toAirport)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(routeInfo); err != nil {
		panic(err)
	}
}

func ShowAirline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airlineCode := strings.Trim(vars["airlineCode"], " ")
	routeInfo := GetRows("routes", "Airline", airlineCode)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(routeInfo); err != nil {
		panic(err)
	}
}
