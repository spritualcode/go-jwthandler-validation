package handler

import (
	"encoding/json"
	"net/http"
)

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	privateKey := "USE_SHARED_KEY"
	accessToken := r.Header.Get("accessToken")
	var errorResponse ErrorResponse

	if len(accessToken) == 0 {
		errorResponse = ErrorResponse{Type: "Invalid Request", Message: "AccessToken Missing from Header"}
		json.NewEncoder(w).Encode(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authFlag, errorToken := ValidateAccessToken(accessToken,privateKey)

	if (errorToken != nil &&  authFlag == false) {
		errorResponse = ErrorResponse{Type: "Invalid Request", Message: "AccessToken for this request is not valid"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

		errorResponse = ErrorResponse{Type: "Success", 	Message: "AccessToken was successfully"}
		json.NewEncoder(w).Encode(errorResponse)
		w.WriteHeader(http.StatusOK)
		return
}

// ErrorResponse Model that represents Error
type ErrorResponse struct {
	Type         string `json:"errorType,omitempty"`
	Message string `json:"errorMessage,omitempty"`
}