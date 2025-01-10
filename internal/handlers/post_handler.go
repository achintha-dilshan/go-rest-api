package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/achintha-dilshan/go-rest-api/internal/models"
	"github.com/achintha-dilshan/go-rest-api/internal/services"
	"github.com/achintha-dilshan/go-rest-api/internal/types"
	"github.com/achintha-dilshan/go-rest-api/internal/utils/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type postHandler struct {
	service services.PostService
}

type PostHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	GetSinglePost(w http.ResponseWriter, r *http.Request)
	EditPost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

func NewPostHandler(service services.PostService) PostHandler {
	return &postHandler{
		service: service,
	}
}

// create a new post
func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title" validate:"required,min=3"`
		Body  string `json:"body" validate:"required,min=3"`
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

	// create new post
	newPost := models.Post{
		AuthorId: int64(userID),
		Title:    req.Title,
		Body:     req.Body,
	}
	postId, err := h.service.CreatePost(r.Context(), &newPost)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		fmt.Println(err)
		return
	}

	// send success response
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"id":        postId,
		"author_id": userID,
		"title":     newPost.Title,
		"body":      newPost.Body,
		"message":   "Post created successfully.",
	})
}

// get all posts
func (h *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.FindAll(r.Context())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"posts": posts,
	})
}

// get single post
func (h *postHandler) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(postId)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid post ID.",
		})
		return
	}

	// find the post
	post, err := h.service.FindPostById(r.Context(), int64(intId))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"post": post,
	})
}

// edit a post
func (h *postHandler) EditPost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title" validate:"required,min=3"`
		Body  string `json:"body" validate:"required,min=3"`
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

	postId := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(postId)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid post ID.",
		})
		return
	}

	// find the post
	post, err := h.service.FindPostById(r.Context(), int64(intId))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	// update post
	post.Title = req.Title
	post.Body = req.Body
	if err := h.service.UpdatePost(r.Context(), post); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Internal server error.",
		})
		return
	}

	// send success response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"title":   post.Title,
		"body":    post.Body,
		"message": "Post updated successfully.",
	})
}

// delete a post
func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(postId)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid post ID.",
		})
		return
	}

	if err := h.service.DeletePost(r.Context(), int64(intId)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": "Post does not exists.",
			})
			return
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": "Internal server error.",
			})
			return
		}
	}

	// send success response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"message": "Post deleted successfully.",
	})
}
