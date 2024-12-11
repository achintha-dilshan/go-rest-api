package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/database"
	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/repositories"
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

func NewAuthHandler() AuthHandler {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository)
	validator := validator.New()

	return &authHandler{
		service:   userService,
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
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate user inputs
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// hash the password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	req.Password = hashedPassword

	// create a new user
	userId, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// send success response
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"user_id": userId,
		"message": "User registered successfully",
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
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate input
	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// retrieve user by email
	user, err := h.service.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// compare passwords
	if err := comparePasswords(user.Password, req.Password); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
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
