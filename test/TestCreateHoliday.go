package test

import (
	"ServiceProductionCalendar/spc"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHoliday(t *testing.T) {

	router := gin.Default()

	//router.POST("/TCH", spc.CreateHoliday)

	w := httptest.NewRecorder()

	sampleHoliday := spc.Holiday{
		Data: "14-04-2023",
	}

	jsonHoliday, _ := json.Marshal(sampleHoliday)

	req, _ := http.NewRequest("POST", "/TCH", bytes.NewBuffer(jsonHoliday))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	expectedResponse, _ := json.Marshal(sampleHoliday)

	if w.Body.String() != string(expectedResponse) {
		t.Fatalf("Expected response body %s, got %s", expectedResponse, w.Body.String())
	}
}
