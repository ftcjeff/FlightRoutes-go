package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8085"
	ready := make(chan bool, 1)

	go func(ready chan bool) {
		router := NewRouter()
		ready <- true

		log.Fatal(http.ListenAndServe(port, router))
	}(ready)

	<-ready

	RegisterSelf("FlightRoutes", "127.0.0.1", port, "route")
	myURI := GetServiceURI("FlightRoutes")
	fmt.Println("\nI'm registered on", myURI)

	fmt.Println("\nListening on port ", port)
	fmt.Print("\nPress any key to stop the server...")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exiting...")
}
