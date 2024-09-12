# DataFlow API Service

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
  "store_id": "string",
  "date": "string",
  "sales": "number"
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
{
  "data": [
    {
      "store_id": "string",
      "date": "string",
      "sales": "number"
    }
  ]
}
```

### 3. Calculate Sales
- **Method**: `POST`
- **Request Body**:
```json
{
  "store_id": "string",
  "start_date": "string",
  "end_date": "string"
}
```
- **Response**:
```json
{
  "total_sales": "number"
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
