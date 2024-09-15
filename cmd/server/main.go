//go:build !testcoverage

package main

import (
	"go.uber.org/zap"
	"net/http"
	"pf_dataflow_api/internal/api"
	"pf_dataflow_api/internal/repository"
	"pf_dataflow_api/internal/service"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := repository.NewInMemorySalesRepository()
	service := &service.SalesService{Repo: repo}
	handler := &api.SalesHandler{Service: service, Logger: logger}

	logger.Info("Starting server on :8080")
	http.HandleFunc("/data", handler.AddSale)
	http.HandleFunc("/data", handler.GetSales)
	http.HandleFunc("/calculate", handler.CalculateTotalSales)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
