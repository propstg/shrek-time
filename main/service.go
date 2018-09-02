package shrek

import (
	"time"
)

const ONE_SHREK float64 = 93.0 * 60.0

func Now() (float64) {
	currentTime := time.Now().UTC()
	return ToShrek(currentTime)
}

func ToShrek(timeToConvert time.Time) (float64) {
	shrepoch := time.Date(2001, time.April, 22, 7, 0, 0, 0, time.UTC)
	secondsSinceShrepoch := timeToConvert.Sub(shrepoch).Seconds()
	return secondsSinceShrepoch / ONE_SHREK
}

func FromShrek(shrekTime float64) (time.Time) {
	shrepoch := time.Date(2001, time.April, 22, 7, 0, 0, 0, time.UTC)
	secondsSinceShrepoch := time.Duration(ONE_SHREK * shrekTime) * time.Second
	return shrepoch.Add(secondsSinceShrepoch)
}
