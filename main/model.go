package main

import "time"

type ShrekResponse struct {
    ShrekStandardTime float64 `json:"sst,omitempty"`
}

type UtcResponse struct {
    UtcValue time.Time `json:"utc,omitempty"`
}
