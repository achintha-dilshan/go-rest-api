package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/achintha-dilshan/go-rest-api/internal/types"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "Authorization header missing.",
			})
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "Invalid Authorization header format.",
			})
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(config.Env.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "Invalid token.",
			})
			return
		}

		// Extract claims and get user ID
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "Invalid token claims.",
			})
			return
		}

		userID, ok := claims["id"].(float64) // JWT encodes numbers as float64
		if !ok {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "User ID missing in token.",
			})
			return
		}

		// Add user ID to the context
		ctx := context.WithValue(r.Context(), types.UserIDKey, int(userID))

		// Pass to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
