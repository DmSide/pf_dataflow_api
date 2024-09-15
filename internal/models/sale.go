package models

import "time"

type Sale struct {
	ProductID string    `json:"product_id" validate:"required"`
	StoreID   string    `json:"store_id" validate:"required"`
	Quantity  int       `json:"quantity_sold" validate:"required,gt=0"`
	Price     float64   `json:"sale_price" validate:"required"` // TODO: Use decimal(github.com/shopspring/decimal) instead of float64
	SaleDate  time.Time `json:"sale_date" validate:"required"`
}

type SalesRepository interface {
	AddSale(sale Sale) error
	GetSales() ([]Sale, error)
	GetSalesByStore(storeID string, startDate, endDate time.Time) (float64, error)
}
