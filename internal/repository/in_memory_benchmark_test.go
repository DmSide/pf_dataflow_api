package repository

import (
	"pf_dataflow_api/internal/models"
	"testing"
	"time"
)

func BenchmarkAddSale(b *testing.B) {
	repo := NewInMemorySalesRepository()
	sale := models.Sale{
		ProductID: "12345",
		StoreID:   "6789",
		Quantity:  10,
		Price:     19.99,
		SaleDate:  time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.AddSale(sale)
	}
}

func BenchmarkGetSales(b *testing.B) {
	repo := NewInMemorySalesRepository()
	for i := 0; i < 1000; i++ {
		repo.AddSale(models.Sale{
			ProductID: "12345",
			StoreID:   "6789",
			Quantity:  10,
			Price:     19.99,
			SaleDate:  time.Now(),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.GetSales()
	}
}
