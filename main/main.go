package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/now", NowController).Methods("GET")
    router.HandleFunc("/toShrek/{utc}", ToShrekController).Methods("GET")
    router.HandleFunc("/fromShrek/{sst}", FromShrekController).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
