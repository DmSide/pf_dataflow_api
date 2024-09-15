package service

import (
	"math"
	"pf_dataflow_api/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockSalesRepository struct {
	sales []models.Sale
}

func (m *MockSalesRepository) AddSale(sale models.Sale) error {
	m.sales = append(m.sales, sale)
	return nil
}

func (m *MockSalesRepository) GetSales() ([]models.Sale, error) {
	return m.sales, nil
}

func (m *MockSalesRepository) GetSalesByStore(storeID string, startDate, endDate time.Time) (float64, error) {
	var totalSales float64
	for _, sale := range m.sales {
		if sale.StoreID == storeID && sale.SaleDate.After(startDate) && sale.SaleDate.Before(endDate) {
			totalSales += math.Round(sale.Price*float64(sale.Quantity)*100) / 100
		}
	}
	return totalSales, nil
}

func TestAddSale(t *testing.T) {
	repo := &MockSalesRepository{}
	service := &SalesService{Repo: repo}

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	err := service.AddSale(sale)
	assert.NoError(t, err, "Error should be nil")

	sales, err := service.GetAllSales()
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, 1, len(sales), "There should be 1 sale")
}

func TestCalculateTotalSales(t *testing.T) {
	repo := &MockSalesRepository{}
	service := &SalesService{Repo: repo}

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	if err := service.AddSale(sale); err != nil {
		t.Errorf("Failed to add sale: %v", err)
	}

	start := time.Now().Add(-time.Hour * 24)
	end := time.Now().Add(time.Hour * 24)

	totalSales, err := service.CalculateTotalSales("1", start, end)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, 199.90, totalSales, "Total sales should be 199.90")
}
