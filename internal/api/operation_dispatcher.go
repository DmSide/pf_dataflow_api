package api

import (
	"net/http"
	"pf_dataflow_api/internal/models"
)

var operationHandlers = map[models.Operation]func(http.ResponseWriter, []byte, *SalesHandler){
	models.TotalSales: handleTotalSales,
}
