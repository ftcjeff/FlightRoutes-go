package main

import (
	"fmt"
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
	dat, err := ioutil.ReadFile("routes.csv")
	check(err)

	ports := strings.Split(string(dat), "\n")

	for _, port := range ports {
		if strings.Contains(port, ",") {
			tokens := strings.Split(port, ",")

			go RepoCreateRoute(
				FlightRoute{Airline: tokens[0],
					AirlineId:         tokens[1],
					SourceAirportCode: tokens[2],
					SourceAirportId:   tokens[3],
					DestAirportCode:   tokens[4],
					DestAirportId:     tokens[5],
					Codeshare:         tokens[6],
					Stops:             tokens[7],
					Equipment:         tokens[8]})
		}
	}
}

func RepoFindRoutes(fromAirport, toAirport string) FlightRoutes {
	var rv FlightRoutes

	for _, t := range flightRoutes {
		if strings.ToLower(t.SourceAirportCode) == strings.ToLower(fromAirport) &&
			strings.ToLower(t.DestAirportCode) == strings.ToLower(toAirport) {
			rv = append(rv, t)
		}
	}

	return rv
}

func RepoCreateRoute(t FlightRoute) FlightRoute {
	fmt.Println("Creating ", t)
	flightRoutes = append(flightRoutes, t)
	return t
}
