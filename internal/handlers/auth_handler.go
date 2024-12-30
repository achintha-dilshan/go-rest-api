package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/achintha-dilshan/go-rest-api/internal/utils/validator"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	service services.UserService
}

type AuthHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(service services.UserService) AuthHandler {
	return &authHandler{
		service: service,
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
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid JSON payload.",
		})
		return
	}

	// validate user inputs
	validator := validator.New()
	if err := validator.Validate(req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	// check if the email is already exist
	exists, err := h.service.ExistUserByEmail(r.Context(), req.Email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	if exists {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, map[string]interface{}{
			"error": map[string]interface{}{
				"email": "Email already exist.",
			},
		})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
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
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	// send success response
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"id":      userId,
		"name":    req.Name,
		"email":   req.Email,
		"message": "User registered successfully.",
	})

}

// login user
func (h *authHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"required,min=3"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid JSON payload.",
		})
		return
	}

	// validate user inputs
	validator := validator.New()
	if err := validator.Validate(req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	// retrieve user by email
	user, err := h.service.FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	if user == nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Email or password is incorrect.",
		})
		return
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Email or password is incorrect.",
		})
		return
	}

	// TODO: generate a token
	token := "mocked_token"

	// send success response
	render.Status(r, http.StatusBadGateway)
	render.JSON(w, r, map[string]interface{}{
		"token":   token,
		"message": "Login successful.",
	})
}
