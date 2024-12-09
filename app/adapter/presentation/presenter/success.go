package presenter

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

func RespondStatusOK[T any](w http.ResponseWriter, respBody T) {
	respondSuccess(w, http.StatusOK, respBody)
}

func respondSuccess[T any](w http.ResponseWriter, statusCode int, respBody T) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(statusCode)
	jsonResp := SuccessResponse[T]{
		Status: statusCode,
		Data:   respBody,
	}
	if err := json.NewEncoder(w).Encode(jsonResp); err != nil {
		RespondStatusInternalServerError(w, err.Error())
	}
}
