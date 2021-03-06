package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/now", NowController).Methods("GET")
	router.HandleFunc("/api/toShrek/{utc}", ToShrekController).Methods("GET")
	router.HandleFunc("/api/fromShrek/{sst}", FromShrekController).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
