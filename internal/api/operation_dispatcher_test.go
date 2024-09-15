package api

import (
	"pf_dataflow_api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationHandlers(t *testing.T) {
	handler, exists := operationHandlers[models.TotalSales]
	assert.True(t, exists, "Handler for TotalSales should exist")
	assert.NotNil(t, handler, "Handler for TotalSales should not be nil")
}
