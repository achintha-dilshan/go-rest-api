package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/achintha-dilshan/go-rest-api/internal/types"
	"github.com/achintha-dilshan/go-rest-api/internal/utils/validator"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	service services.UserService
}

type UserHandler interface {
	ResetPassword(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

// reset password
func (h *userHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OldPassword        string `json:"old_password" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=3"`
		ConfirmNewPassword string `json:"confirm_new_password"`
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

	// get user ID from the context
	userID, ok := r.Context().Value(types.UserIDKey).(int)
	if !ok {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "Ensure that you are logged in.",
		})
		return
	}

	// retrieve user by id
	user, err := h.service.FindUserById(r.Context(), int64(userID))
	if err != nil || user == nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "User not found or unauthorized.",
		})
		return
	}

	// compare old password with stored password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Old password is incorrect.",
		})
		return
	}

	// ensure the new password is different
	if req.OldPassword == req.NewPassword {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "New password should be different from the old password.",
		})
		return
	}

	// compare new password and confirmation password
	if req.NewPassword != req.ConfirmNewPassword {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "New password and confirmation password do not match.",
		})
		return
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to hash the new password.",
		})
		return
	}

	// update user's password
	user.Password = string(hashedPassword)
	if err := h.service.UpdateUser(r.Context(), user); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to update the password.",
		})
		return
	}

	// send success response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "Password reset was successful.",
	})
}

// update user
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name" validate:"required,min=3"`
		Email string `json:"email" validate:"required,email"`
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

	// get user ID from the context
	userID, ok := r.Context().Value(types.UserIDKey).(int)
	if !ok {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "Ensure that you are logged in.",
		})
		return
	}

	// retrieve user by id
	user, err := h.service.FindUserById(r.Context(), int64(userID))
	if err != nil || user == nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "User not found or unauthorized.",
		})
		return
	}

	// updated user
	user.Name = req.Name
	user.Email = req.Email
	if err := h.service.UpdateUser(r.Context(), user); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	// send success response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"name":    user.Name,
		"email":   user.Email,
		"message": "User updated successfully.",
	})
}

// delete user
func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get user ID from the context
	userID, ok := r.Context().Value(types.UserIDKey).(int)
	if !ok {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "Ensure that you are logged in.",
		})
		return
	}

	// retrieve user by id
	user, err := h.service.FindUserById(r.Context(), int64(userID))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	if user == nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "Ensure that you are logged in.",
		})
		return
	}

	if err := h.service.DeleteUser(r.Context(), user.Id); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	// send success response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "User deleted successfully.",
	})
}
