package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// init router
	router := mux.NewRouter()
	// endpoints
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.Handle).Methods(route.Method)
	}
	// router.HandleFunc("/getLatest", getLatest).Methods("GET")
	// router.HandleFunc("/update", test).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

