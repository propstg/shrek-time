package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func NowController(w http.ResponseWriter, r *http.Request) {
	writeShrekResponse(w, Now())
}

func ToShrekController(w http.ResponseWriter, r *http.Request) {
	utc, err := time.Parse(time.RFC3339, mux.Vars(r)["utc"])
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
	sst, err := strconv.ParseFloat(mux.Vars(r)["sst"], 64)
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
