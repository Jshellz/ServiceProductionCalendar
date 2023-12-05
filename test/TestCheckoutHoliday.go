package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckoutHoliday(t *testing.T) {

	router := gin.Default()

	_ = []string{
		"01-01-2022",
		"25-12-2022",
	}
	//router.GET("/TCH", spc.CheckoutHoliday)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PATCH", "/holidays?id=1", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"data":"0"}`

	if w.Body.String() != expectedResponse {
		t.Fatalf("Expected response body %s, got %s", expectedResponse, w.Body.String())
	}
}
