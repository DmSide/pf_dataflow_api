package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
}

func NewValidationErrorResponse(err error) Response {
	return Response{
		Status:  "error",
		Message: "Validation error: " + err.Error(),
	}
}

func NewInternalErrorResponse(err error) Response {
	return Response{
		Status:  "error",
		Message: "Internal error: " + err.Error(),
	}
}

func NewDataDecodeErrorResponse(err error) Response {
	return Response{
		Status:  "error",
		Message: "Data decoder error: " + err.Error(),
	}
}

func NewOkResponse() Response {
	return Response{
		Status: "success",
	}
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
