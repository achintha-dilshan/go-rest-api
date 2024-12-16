package utils

import (
	"fmt"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/pkg/res"
	"github.com/go-playground/validator/v10"
)

// sending validator errors as JSON
func HandleValidationErrors(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]interface{})
		for _, ve := range validationErrors {
			errors[ve.Field()] = fmt.Sprintf("failed on the '%s' tag", ve.Tag())
		}

		// Use RespondWithJSON
		res.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	// Fallback for non-validation errors
	res.JSON(w, http.StatusBadRequest, map[string]interface{}{
		"error": "invalid request",
	})
}
