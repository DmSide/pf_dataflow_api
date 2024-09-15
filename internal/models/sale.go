package models

import "time"

type Sale struct {
	ProductID string    `json:"product_id"`
	StoreID   string    `json:"store_id"`
	Quantity  int       `json:"quantity_sold"`
	Price     float64   `json:"sale_price"` // TODO: Use decimal(github.com/shopspring/decimal) instead of float64
	SaleDate  time.Time `json:"sale_date"`
}

type SalesRepository interface {
	AddSale(sale Sale) error
	GetSales() ([]Sale, error)
	GetSalesByStore(storeID string, startDate, endDate time.Time) (float64, error)
}
