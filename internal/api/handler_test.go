package api

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"pf_dataflow_api/internal/models"
	"pf_dataflow_api/internal/repository"
	"pf_dataflow_api/internal/service"
	"strings"
	"testing"
	"time"
)

func TestAddSale(t *testing.T) {
	repo := repository.NewInMemorySalesRepository()
	salesService := &service.SalesService{Repo: repo}
	logger, _ := zap.NewDevelopment()
	handler := &SalesHandler{Service: salesService, Logger: logger}

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	body, _ := json.Marshal(sale)
	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.AddSale(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"success"}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCalculateTotalSales(t *testing.T) {
	repo := repository.NewInMemorySalesRepository()
	salesService := &service.SalesService{Repo: repo}
	logger, _ := zap.NewDevelopment()
	handler := &SalesHandler{Service: salesService, Logger: logger}

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	if err := salesService.AddSale(sale); err != nil {
		t.Errorf("Failed to add sale: %v", err)
	}

	requestBody := map[string]string{
		"operation":  "total_sales",
		"store_id":   "1",
		"start_date": time.Now().Add(-time.Hour * 24).Format(time.RFC3339),
		"end_date":   time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CalculateTotalSales(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil || response["total_sales"].(float64) != 199.90 {
		t.Errorf("Expected total_sales 199.90, got %v", response["total_sales"])
	}
}

func TestGetSales(t *testing.T) {
	repo := repository.NewInMemorySalesRepository()
	salesService := &service.SalesService{Repo: repo}
	logger, _ := zap.NewDevelopment()
	handler := &SalesHandler{Service: salesService, Logger: logger}

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}
	if err := salesService.AddSale(sale); err != nil {
		t.Errorf("Failed to add sale: %v", err)

	}

	req, _ := http.NewRequest("GET", "/data", nil)
	rr := httptest.NewRecorder()

	handler.GetSales(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var sales []models.Sale
	err := json.NewDecoder(rr.Body).Decode(&sales)
	if err != nil || len(sales) != 1 {
		t.Errorf("Expected 1 sale, got %v", len(sales))
	}
}
