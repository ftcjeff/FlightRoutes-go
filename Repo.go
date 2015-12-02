package main

import (
	"io/ioutil"
	"strings"
)

var flightRoutes FlightRoutes

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	tableName := "routes"
	tableFields := []string{"Airline", "AirlineId", "SourceAirportCode", "SourceAirportId", "DestAirportCode", "DestAirportId", "Codeshare", "Stops", "Equipment"}

	mysql := GetServiceURI("mysql")
	CreateTable(mysql, "picasso", "picasso", "picasso", tableName)

	dat, err := ioutil.ReadFile("routes.csv")
	check(err)

	ports := strings.Split(string(dat), "\n")

	for _, port := range ports {
		if strings.Contains(port, ",") {
			tokens := strings.Split(port, ",")

			AddRow(tableName, tableFields, tokens)
		}
	}
}
