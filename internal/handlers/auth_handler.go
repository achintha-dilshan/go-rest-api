package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/achintha-dilshan/go-rest-api/internal/utils"
	"github.com/achintha-dilshan/go-rest-api/pkg/res"
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

// register user
func (h *authHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid JSON payload.",
			"message": "Invalid JSON payload.",
		})
		return
	}

	// validate user inputs
	if err := h.validator.Struct(req); err != nil {
		utils.HandleValidationErrors(w, err)
		return
	}

	// check if the email is already exist
	emailExist, err := h.service.ExistUserByEmail(r.Context(), req.Email)
	if err != nil {
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal server error.",
			"message": "Internal server error.",
		})
		return
	}

	if emailExist {
		res.JSON(w, http.StatusConflict, map[string]interface{}{
			"error": map[string]interface{}{
				"email": "Email already exist.",
			},
			"message": "Email already exist.",
		})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal server error.",
			"message": "Internal server error.",
		})
		return
	}
	req.Password = string(hashedPassword)

	// create a new user
	newUser := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userId, err := h.service.CreateUser(r.Context(), &newUser)
	if err != nil {
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal server error.",
			"message": "Internal server error.",
		})
		return
	}

	// send success response
	res.JSON(w, http.StatusCreated, map[string]interface{}{
		"id":      userId,
		"name":    req.Name,
		"email":   req.Email,
		"message": "User registered successfully.",
	})
}

// login user
func (h *authHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid JSON payload.",
			"message": "Invalid JSON payload.",
		})
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
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal server error.",
			"message": "Internal server error.",
		})
		return
	}

	if user == nil {
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Email or password is incorrect.",
			"message": "Email or password is incorrect.",
		})
		return
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		res.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   "Email or password is incorrect.",
			"message": "Email or password is incorrect.",
		})
		return
	}

	// TODO: generate a token
	token := "mocked_token"

	// send success response
	res.JSON(w, http.StatusOK, map[string]interface{}{
		"token":   token,
		"message": "Login successful.",
	})
}
