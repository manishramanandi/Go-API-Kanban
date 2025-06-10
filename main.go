package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.Handle("/register", alice.New(loggingMiddleware).ThenFunc(register)).Methods("POST")

	router.Handle("/login", alice.New(loggingMiddleware).ThenFunc(login)).Methods("POST")

	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(createProject)).Methods("POST")
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(updateProjects)).Methods("PUT")

	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(getProjects)).Methods("GET")

	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(getProject)).Methods("GET")
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(deletProject)).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// register user
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "From register "})
}

// login
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "From login "})
}

// create projects
func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from createProject"})
}

// getproject
func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from updateProject"})
	ID: id,
}

// delete project
func deletProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "from deletProject"})
	ID: id,
}
