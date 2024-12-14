package presenter

import (
	"encoding/json"
	"net/http"
)

type FailureResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondInternalServerError(w http.ResponseWriter, message string) {
	respondJsonFailure(w, http.StatusInternalServerError, message)
}

func respondJsonFailure(w http.ResponseWriter, statusCode int, message string) {
	jsonResp := FailureResponse{
		Status:  statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(jsonResp)
}
