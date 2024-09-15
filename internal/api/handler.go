package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io"
	"net/http"
	"pf_dataflow_api/internal/models"
	"pf_dataflow_api/internal/service"
)

type SalesHandler struct {
	Service  *service.SalesService
	Logger   *zap.Logger
	Validate *validator.Validate
}

func NewSalesHandler(service *service.SalesService, logger *zap.Logger) *SalesHandler {
	return &SalesHandler{
		Service:  service,
		Logger:   logger,
		Validate: validator.New(),
	}
}

func (h *SalesHandler) AddSale(w http.ResponseWriter, r *http.Request) {
	var sale models.Sale
	if err := json.NewDecoder(r.Body).Decode(&sale); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		response := NewDataDecodeErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if err := h.Validate.Struct(sale); err != nil {
		h.Logger.Error("Validation failed", zap.Error(err))
		response := NewValidationErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if err := h.Service.AddSale(sale); err != nil {
		h.Logger.Error("Failed to add sale", zap.Error(err))
		response := NewInternalErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	h.Logger.Info("Sale added successfully", zap.String("product_id", sale.ProductID), zap.String("store_id", sale.StoreID))
	WriteJSONResponse(w, http.StatusOK, NewOkResponse())
}

func (h *SalesHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.Service.GetAllSales()
	if err != nil {
		h.Logger.Error("Failed to retrieve sales", zap.Error(err))
		response := NewInternalErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	h.Logger.Info("Retrieved all sales", zap.Int("total_sales", len(sales)))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}

func (h *SalesHandler) CalculateTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error("Failed to read body", zap.Error(err))
		response := NewDataDecodeErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var req struct {
		Operation string `json:"operation" validate:"required"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		h.Logger.Error("Invalid input", zap.Error(err))
		response := NewValidationErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		h.Logger.Error("Validation failed", zap.Error(err))
		response := NewValidationErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if !models.IsValidOperation(models.Operation(req.Operation)) {
		h.Logger.Error("Invalid operation")
		response := NewValidationErrorResponse(err)
		WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	operationHandlers[models.Operation(req.Operation)](w, body, h)
}
