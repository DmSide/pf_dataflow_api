package models

type Operation string

const (
	TotalSales Operation = "total_sales"
)

var allowedOperations = []Operation{
	TotalSales,
	// Add more if you need
}

func IsValidOperation(op Operation) bool {
	for _, validOp := range allowedOperations {
		if op == validOp {
			return true
		}
	}
	return false
}
