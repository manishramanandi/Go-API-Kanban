package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteResponse struct {
	Message string `json: "message"`
}

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/projects", createProject).Methods("POST")
	router.HandleFunc("/projects{id}", updateProjects).Methods("PUT")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/projects{id}", getProject).Methods("GET")
	router.HandleFunc("/projects{id}", deletProject).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "From register "})
}

// login

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "From register "})
}

// create projects
func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from createProject"})
}

// getproject
func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from from getproject"})
}

// getprojects
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from getProjects"})
}

// update projects
func updateProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from updateProject"})
}

// delete project
func deletProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from deletProject"})
}
