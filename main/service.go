package main

import "time"

const ONE_SHREK float64 = 93.0 * 60.0

func Now() float64 {
    return ToShrek(time.Now().UTC())
}

func ToShrek(timeToConvert time.Time) float64 {
    return timeToConvert.Sub(getShrepoch()).Seconds() / ONE_SHREK
}

func FromShrek(shrekTime float64) time.Time {
    return getShrepoch().Add(time.Duration(ONE_SHREK * shrekTime) * time.Second)
}

func getShrepoch() time.Time {
    return time.Date(2001, time.April, 22, 7, 0, 0, 0, time.UTC)
}
