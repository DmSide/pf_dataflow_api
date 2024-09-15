package models

type Operation string

const (
	TotalSales Operation = "total_sales"
)

var allowedOperations = []Operation{
	TotalSales,
}

func IsValidOperation(op Operation) bool {
	for _, validOp := range allowedOperations {
		if op == validOp {
			return true
		}
	}
	return false
}
