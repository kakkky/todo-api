package presenter

import (
	"encoding/json"
	"net/http"
)

type FailureResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondBadRequest(rw http.ResponseWriter, message string) {
	respondJsonFailure(rw, http.StatusBadRequest, message)
}

func RespondInternalServerError(rw http.ResponseWriter, message string) {
	respondJsonFailure(rw, http.StatusInternalServerError, message)
}

func RespondUnAuthorized(rw http.ResponseWriter, message string) {
	respondJsonFailure(rw, http.StatusUnauthorized, message)
}

func respondJsonFailure(rw http.ResponseWriter, statusCode int, message string) {
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	rw.WriteHeader(statusCode)
	jsonResp := FailureResponse{
		Status:  statusCode,
		Message: message,
	}
	json.NewEncoder(rw).Encode(jsonResp)
}
