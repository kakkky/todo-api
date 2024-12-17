package presenter

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

func RespondOK[T any](w http.ResponseWriter, respBody T) {
	respondJsonSuccess(w, http.StatusOK, respBody)
}

func RespondCreated[T any](w http.ResponseWriter, respBody T) {
	respondJsonSuccess(w, http.StatusCreated, respBody)
}

func respondJsonSuccess[T any](w http.ResponseWriter, statusCode int, respBody T) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(statusCode)
	jsonResp := SuccessResponse[T]{
		Status: statusCode,
		Data:   respBody,
	}
	if err := json.NewEncoder(w).Encode(jsonResp); err != nil {
		RespondInternalServerError(w, err.Error())
	}
}
