package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/achintha-dilshan/go-rest-api/internal/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	service   services.UserService
	validator *validator.Validate
}

type AuthHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(service services.UserService) AuthHandler {
	validator := validator.New()
	return &authHandler{
		service:   service,
		validator: validator,
	}
}

// hash password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// compare passwords
func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// register user
func (h *authHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.User

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload. Please check your input."})
		return
	}

	// validate user inputs
	if err := req.Validate(); err != nil {
		utils.HandleValidationErrors(w, err)
		return
	}

	// hash the password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to hash password. Please try again later."})
		return
	}
	req.Password = hashedPassword

	// create a new user
	userId, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"error": "A user with this email already exists. Please use a different email."})
		return
	}

	// send success response
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"user_id": userId,
		"message": "User registered successfully",
	})
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// login user
func (h *authHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload. Please check your input."})
		return
	}

	// validate input
	if err := h.validator.Struct(req); err != nil {
		utils.HandleValidationErrors(w, err)
		return
	}

	// retrieve user by email
	user, err := h.service.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Email or password is incorrect. Please try again."})
		return
	}

	// compare passwords
	if err := comparePasswords(user.Password, req.Password); err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Email or password is incorrect. Please try again."})
		return
	}

	// TODO: generate a token
	token := "mocked_token"

	// send success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"token":   token,
		"message": "Login successful",
	})
}
