//go:build concurrency

package repository

import (
	"github.com/stretchr/testify/assert"
	"pf_dataflow_api/internal/models"
	"sync"
	"testing"
	"time"
)

func TestInMemorySalesRepositoryConcurrency(t *testing.T) {
	repo := NewInMemorySalesRepository()
	var wg sync.WaitGroup
	storeID := "store_1"
	startDate := time.Now().Add(-time.Hour * 24)
	endDate := time.Now().Add(time.Hour * 24)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sale := models.Sale{
				ProductID: "product_1",
				StoreID:   storeID,
				Quantity:  10,
				Price:     19.99,
				SaleDate:  time.Now(),
			}
			err := repo.AddSale(sale)
			assert.NoError(t, err, "AddSale should not return an error")
		}(i)
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			totalSales, err := repo.GetSalesByStore(storeID, startDate, endDate)
			assert.NoError(t, err, "GetSalesByStore should not return an error")
			assert.GreaterOrEqual(t, totalSales, 0.0, "Total sales should be >= 0")
		}()
	}

	wg.Wait()

	totalSales, err := repo.GetSalesByStore(storeID, startDate, endDate)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, totalSales, 19990.0, "Total sales should be at least 19990")
}
