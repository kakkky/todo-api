package presenter

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

func RespondOK[T any](rw http.ResponseWriter, respBody T) {
	respondJsonSuccess(rw, http.StatusOK, respBody)
}

func RespondCreated[T any](rw http.ResponseWriter, respBody T) {
	respondJsonSuccess(rw, http.StatusCreated, respBody)
}

func RespondNoContent(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}

func respondJsonSuccess[T any](rw http.ResponseWriter, statusCode int, respBody T) {
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	rw.WriteHeader(statusCode)
	jsonResp := SuccessResponse[T]{
		Status: statusCode,
		Data:   respBody,
	}
	if err := json.NewEncoder(rw).Encode(jsonResp); err != nil {
		RespondInternalServerError(rw, err.Error())
	}
}
