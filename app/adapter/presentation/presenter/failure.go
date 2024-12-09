package presenter

import (
	"encoding/json"
	"net/http"
)

type FailureResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondStatusInternalServerError(w http.ResponseWriter, message string) {
	respondFailure(w, http.StatusInternalServerError, message)
}

func respondFailure(w http.ResponseWriter, statusCode int, message string) {
	jsonResp := FailureResponse{
		Status:  statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(jsonResp)
}
