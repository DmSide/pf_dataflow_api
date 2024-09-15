package service

import (
	"pf_dataflow_api/internal/models"
	"time"
)

type SalesService struct {
	Repo models.SalesRepository
}

func (s *SalesService) AddSale(sale models.Sale) error {
	return s.Repo.AddSale(sale)
}

func (s *SalesService) GetAllSales() ([]models.Sale, error) {
	return s.Repo.GetSales()
}

func (s *SalesService) CalculateTotalSales(storeID string, startDate, endDate time.Time) (float64, error) {
	return s.Repo.GetSalesByStore(storeID, startDate, endDate)
}
