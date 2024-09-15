# DataFlow API Service
## How Business Requirements Are Addressed
### Architecture for Scalability and Maintainability

A hexagonal architecture (also known as ports and adapters) is used to ensure modularity and component flexibility. 

In this solution, it's easy to replace the InMemorySalesRepository with another implementation (for example, database). 

The logic is separated from infrastructure, enhancing maintainability and testability.

The logic of calculate sales is managed to work with different models, so it's easy to add new operations in the future and check different type. But for REST API it's better to split it on the different endpoint. It also will be easier to test different scenarios.

### Performance

To ensure safe concurrent access, sync.RWMutex is used. This provides high performance during read operations (allowing multiple simultaneous reads) and ensures safe writes.

### Data Storage

The service uses a map to store sales data. The map key is a composite key of store_id and date and use Mutex to ensure that the same data is not added twice.

## Overview
This API service is developed for DataFlow Inc. to manage and analyze sales data. The service provides three endpoints:
1. **POST /data** - Add new sales data.
2. **GET /data** - Retrieve all stored sales data.
3. **POST /calculate** - Calculate total sales for a given store within a specified date range.

## Endpoints (Could be moved to Swagger later)
### 1. Add Data
- **Method**: `POST`
- **Request Body**:
```json
{
  "product_id": "12345",
  "store_id": "6789",
  "quantity_sold": 10,
  "sale_price": 19.99,
  "sale_date": "2024-06-15T14:30:00Z"
}
```
- **Response**:
```json
{
  "message": "Data added successfully"
}
```

### 2. Get Data
- **Method**: `GET`
- **Response**:
```json
[
  {
    "product_id": "12345",
    "store_id": "6789",
    "quantity_sold": 10,
    "sale_price": 19.99,
    "sale_date": "2024-06-15T14:30:00Z"
  }
]
```

### 3. Calculate Sales
- **Method**: `POST`
- **Request Body**:
```json
{
  "operation": "total_sales",
  "store_id": "6789",
  "start_date": "2024-06-01T00:00:00Z",
  "end_date": "2024-06-30T23:59:59Z"
}
```
- **Response**:
```json
{
  "store_id": "6789",
  "total_sales": 199.90,
  "start_date": "2024-06-01T00:00:00Z",
  "end_date": "2024-06-30T23:59:59Z"
}
```

## Setup
1. Clone the repository.
2. Install Go 1.18+.
3. Install dependencies using `go mod tidy`.

## Run
1. Run the server using `go run main.go`.
2. The server will start on `localhost:8080`.

## Test
1. Run tests using `go test ./...`.
2. For coverage run `go test -tags=testcoverage -cover ./...`.
3. For running test of concurrency use `go test -race -tags=concurrency ./...`.
