package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pf_dataflow_api/internal/repository"
	"pf_dataflow_api/internal/service"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestHandleTotalSales(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	repo := repository.NewInMemorySalesRepository()
	salesService := &service.SalesService{Repo: repo}
	handler := &SalesHandler{Logger: logger, Service: salesService}

	reqBody := map[string]string{
		"operation":  "total_sales",
		"store_id":   "1",
		"start_date": time.Now().Add(-time.Hour * 24).Format(time.RFC3339),
		"end_date":   time.Now().Format(time.RFC3339),
	}
	body, _ := json.Marshal(reqBody)

	httptest.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handleTotalSales(rr, body, handler)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
