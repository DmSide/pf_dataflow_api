//go:build !testcoverage

package main

import (
	"github.com/gorilla/mux"
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
	salesService := &service.SalesService{Repo: repo}
	handler := api.NewSalesHandler(salesService, logger)

	r := mux.NewRouter()

	r.HandleFunc("/data", handler.AddSale).Methods("POST")
	r.HandleFunc("/data", handler.GetSales).Methods("GET")
	r.HandleFunc("/calculate", handler.CalculateTotalSales).Methods("POST")

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
