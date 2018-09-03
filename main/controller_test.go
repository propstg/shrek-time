package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
    "time"
)

const NOW_URL string = "/api/now"
const FROM_SHREK_TEMPLATE string = "/api/fromShrek/{sst}"
const TO_SHREK_TEMPLATE string = "/api/toShrek/{utc}"

func TestNowController(t *testing.T) {
    responseRecorder := setupTest(t, NOW_URL, NOW_URL, NowController)
    expectSuccessfulResponseCode(t, responseRecorder)
    if response := getShrekResponse(t, responseRecorder); response.ShrekStandardTime < 98211 {
        t.Errorf("sst too low in response: got %v wanted >= 98211", response.ShrekStandardTime)
    }
}

func TestToShrekControllerValidInput(t *testing.T) {
    responseRecorder := setupTest(t, "/api/toShrek/2018-12-27T15:00:00Z", TO_SHREK_TEMPLATE, ToShrekController)
    expectSuccessfulResponseCode(t, responseRecorder)
    if response := getShrekResponse(t, responseRecorder); response.ShrekStandardTime != 100000 {
        t.Errorf("sst response incorrect: got %v wanted 100000", response.ShrekStandardTime)
    }
}

func TestToShrekControllerInvalidInput(t *testing.T) {
    responseRecorder := setupTest(t, "/api/toShrek/invalid", TO_SHREK_TEMPLATE, ToShrekController)
    expectBadRequestResponseCode(t, responseRecorder)
}

func TestFromShrekControllerValidInput(t *testing.T) {
    responseRecorder := setupTest(t, "/api/fromShrek/100000", FROM_SHREK_TEMPLATE, FromShrekController)
    expectSuccessfulResponseCode(t, responseRecorder)

    var response UtcResponse
    if err := json.Unmarshal([]byte(responseRecorder.Body.String()), &response); err != nil {
        t.Fatal(err);
    } else if response.UtcValue != time.Date(2018, time.December, 27, 15, 0, 0, 0, time.UTC) {
        t.Errorf("sst response incorrect: got %v wanted 2018-12-27T15:00:00Z", response.UtcValue)
    }
}

func TestFromShrekControllerInvalidInput(t *testing.T) {
    responseRecorder := setupTest(t, "/api/fromShrek/invalid", FROM_SHREK_TEMPLATE, FromShrekController)
    expectBadRequestResponseCode(t, responseRecorder)
}

func setupTest(t *testing.T, url string, urlPattern string, controller  func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Fatal(err)
    }
    responseRecorder := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc(urlPattern, controller)
    router.ServeHTTP(responseRecorder, req)
    return responseRecorder
}

func expectSuccessfulResponseCode(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
    expectResponseCode(t, responseRecorder, http.StatusOK)
}

func expectBadRequestResponseCode(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
    expectResponseCode(t, responseRecorder, http.StatusBadRequest)
}

func expectResponseCode(t *testing.T, responseRecorder *httptest.ResponseRecorder, expectedStatus int) {
    if status := responseRecorder.Code; status != expectedStatus {
        t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
    }
}

func getShrekResponse(t *testing.T, responseRecorder *httptest.ResponseRecorder) ShrekResponse {
    var response ShrekResponse
    if err := json.Unmarshal([]byte(responseRecorder.Body.String()), &response); err != nil {
        t.Fatal(err);
    }
    return response
}
