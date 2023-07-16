package auth

import (
	"context"
	"log"
	"megachasma/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type userIdKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, req)
			return
		}

		t, err := utils.ValidateJwt(token)
		if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
			userId := (*claims)["sub"].(string)
			ctx := context.WithValue(req.Context(), userIdKey{}, userId)
			next.ServeHTTP(w, req.WithContext(ctx))
		} else {
			next.ServeHTTP(w, req)
			return
		}
		if err != nil {
			log.Println(err)
			http.Error(w, `{"reason": "invalid token"}`, http.StatusUnauthorized)
			return
		}
	})
}
