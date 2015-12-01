package main

type FlightRoute struct {
	Airline           string `json:"airline"`
	AirlineId         string `json:"airline_id"`
	SourceAirportCode string `json:"source_airport_code"`
	SourceAirportId   string `json:"source_airport_id"`
	DestAirportCode   string `json:"dest_airport_code"`
	DestAirportId     string `json:"dest_airport_id"`
	Codeshare         string `json:"codeshare"`
	Stops             string `json:"stops"`
	Equipment         string `json:"equipment"`
}

type FlightRoutes []FlightRoute
