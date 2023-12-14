package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/", Index).Methods("GET")
	route.HandleFunc("/health", HealthCheck).Methods("GET")

	log.Println("Server started on: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", route))
}

func Index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}

// HealthCheck is a function to check the health of the application
func HealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}
