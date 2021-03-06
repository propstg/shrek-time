package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	assert.True(t, float64(98211) < Now())
}

func TestToShrek(t *testing.T) {
	assert.Equal(t, float64(100000), ToShrek(time.Date(2018, time.December, 27, 15, 0, 0, 0, time.UTC)))
	assert.Equal(t, float64(0), ToShrek(time.Date(2001, time.April, 22, 7, 0, 0, 0, time.UTC)))
}

func TestFromShrek(t *testing.T) {
	assert.Equal(t, FromShrek(float64(100000)), time.Date(2018, time.December, 27, 15, 0, 0, 0, time.UTC))
	assert.Equal(t, FromShrek(float64(0)), time.Date(2001, time.April, 22, 7, 0, 0, 0, time.UTC))
}
