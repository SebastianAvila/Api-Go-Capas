// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRouteHandler)
	router.HandleFunc("/task", getTaskHandler).Methods("GET")
	router.HandleFunc("/task/{id}", getOneTaskHandler).Methods("GET")
	router.HandleFunc("/task", createTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", updateTaskHandler).Methods("PUT")
	router.HandleFunc("/task/{id}", deleteTaskHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3089", router))
}
