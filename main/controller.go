package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "strconv"
)

func NowController(w http.ResponseWriter, r *http.Request) {
    writeShrekResponse(w, Now())
}

func ToShrekController(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    utc, err := time.Parse(time.RFC3339, params["utc"])
    if err != nil {
        writeError(w, err)
        return
    }

    writeShrekResponse(w, ToShrek(utc))
}

func writeShrekResponse(w http.ResponseWriter, shrekStandardTime float64) {
    var response ShrekResponse
    response.ShrekStandardTime = shrekStandardTime
    json.NewEncoder(w).Encode(response)
}

func FromShrekController(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    sst, err := strconv.ParseFloat(params["sst"], 64)
    if err != nil {
        writeError(w, err)
        return
    }

    var response UtcResponse
    response.UtcValue = FromShrek(sst)
    json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(err)
}
