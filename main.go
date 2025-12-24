package main

import (
	"4inrow/internal/handlers"

	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	port := ":9080"

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ws/live", handlers.LiveWebSocketHandler)
	http.HandleFunc("/ws/lobby", handlers.LobbyWebSocketHandler)

	log.Println("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, enableCORS(http.DefaultServeMux)))
}
   