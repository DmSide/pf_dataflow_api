package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"pf_dataflow_api/internal/models"
	"pf_dataflow_api/internal/service"
)

type SalesHandler struct {
	Service *service.SalesService
	Logger  *zap.Logger
}

func (h *SalesHandler) AddSale(w http.ResponseWriter, r *http.Request) {
	var sale models.Sale
	if err := json.NewDecoder(r.Body).Decode(&sale); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.AddSale(sale); err != nil {
		h.Logger.Error("Failed to add sale", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("Sale added successfully", zap.String("product_id", sale.ProductID), zap.String("store_id", sale.StoreID))

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *SalesHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.Service.GetAllSales()
	if err != nil {
		h.Logger.Error("Failed to retrieve sales", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("Retrieved all sales", zap.Int("total_sales", len(sales)))

	json.NewEncoder(w).Encode(sales)
}

func (h *SalesHandler) CalculateTotalSales(w http.ResponseWriter, r *http.Request) {
	// Here we save to the buffer to read it one more time inside the function
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error("Failed to read body", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var req struct {
		Operation string `json:"operation"`
		// TODO: here we have to add the store_id, start_date, and end_date fields. It depends on the other operations
	}

	if err := json.Unmarshal(body, &req); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !models.IsValidOperation(models.Operation(req.Operation)) {
		h.Logger.Error("Invalid operation")
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	operationHandlers[models.Operation(req.Operation)](w, body, h)
}
