package handlers

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}
