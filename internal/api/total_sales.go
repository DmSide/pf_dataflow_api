package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func handleTotalSales(w http.ResponseWriter, body []byte, h *SalesHandler) {
	var req struct {
		Operation string `json:"operation"`
		StoreID   string `json:"store_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	start, _ := time.Parse(time.RFC3339, req.StartDate)
	end, _ := time.Parse(time.RFC3339, req.EndDate)

	totalSales, err := h.Service.CalculateTotalSales(req.StoreID, start, end)
	if err != nil {
		h.Logger.Error("Failed to calculate total sales", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Calculated total sales", zap.String("store_id", req.StoreID), zap.Float64("total_sales", totalSales))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"store_id":    req.StoreID,
		"total_sales": totalSales,
		"start_date":  req.StartDate,
		"end_date":    req.EndDate,
	})
}
