package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User logged in successfully"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request Payload")
		return
	}
}
