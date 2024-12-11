package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// sending a JSON response with a given status code and payload
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// sending validator errors as JSON
func HandleValidationErrors(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, ve := range validationErrors {
			errors[ve.Field()] = fmt.Sprintf("failed on the '%s' tag", ve.Tag())
		}

		// Use RespondWithJSON
		RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	// Fallback for non-validation errors
	RespondWithJSON(w, http.StatusBadRequest, map[string]string{
		"error": "invalid request",
	})
}
