package repository_test

import (
	"pf_dataflow_api/internal/models"
	"pf_dataflow_api/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddSale(t *testing.T) {
	repo := repository.NewInMemorySalesRepository()

	sale := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	err := repo.AddSale(sale)
	assert.NoError(t, err)

	sales, err := repo.GetSales()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sales))
}

func TestGetSalesByStore(t *testing.T) {
	repo := repository.NewInMemorySalesRepository()

	now := time.Now()
	sale1 := models.Sale{
		ProductID: "1",
		StoreID:   "1",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  now.Add(-time.Hour * 24),
	}
	sale2 := models.Sale{
		ProductID: "2",
		StoreID:   "2",
		Quantity:  5,
		Price:     9.99,
		SaleDate:  now,
	}

	repo.AddSale(sale1)
	repo.AddSale(sale2)

	startDate := now.Add(-time.Hour * 48)
	endDate := now.Add(time.Hour * 24)

	totalSales, err := repo.GetSalesByStore("1", startDate, endDate)
	assert.NoError(t, err)
	assert.Equal(t, 199.9, totalSales)
}
