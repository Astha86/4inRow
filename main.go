package main

import (
	"4inrow/internal/handlers"

	"log"
	"net/http"
)

func main() {

	port := ":9080"

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ws/live", handlers.LiveWebSocketHandler)
	http.HandleFunc("/ws/lobby", handlers.LobbyWebSocketHandler)

	log.Println("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
