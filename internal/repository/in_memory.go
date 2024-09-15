package repository

import (
	"math"
	"pf_dataflow_api/internal/models"
	"sync"
	"time"
)

type InMemorySalesRepository struct {
	sales []models.Sale
	mu    sync.RWMutex
}

func NewInMemorySalesRepository() *InMemorySalesRepository {
	return &InMemorySalesRepository{
		sales: []models.Sale{},
	}
}

func (r *InMemorySalesRepository) AddSale(sale models.Sale) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sales = append(r.sales, sale)
	return nil
}

func (r *InMemorySalesRepository) GetSales() ([]models.Sale, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.sales, nil
}

func (r *InMemorySalesRepository) GetSalesByStore(storeID string, startDate, endDate time.Time) (float64, error) {
	var totalSales float64
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, sale := range r.sales {
		if sale.StoreID == storeID && sale.SaleDate.After(startDate) && sale.SaleDate.Before(endDate) {
			totalSales += math.Round(sale.Price*float64(sale.Quantity)*100) / 100
		}
	}
	return totalSales, nil
}
