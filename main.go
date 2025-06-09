package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type RouteResponse struct {
	Message string `json: "message"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from the server!"})
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":4269", router))
	fmt.Println("Server is running on port 4269")
}
