package auth

import (
	"context"
	"log"
	"megachasma/utils"
	"net/http"
	"strings"

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
		splitToken := strings.Split(token, " ")
		if len(splitToken) == 2 {
			if splitToken[0] == "Bearer" {
				t, err := utils.ValidateJwt(splitToken[1])
				if err != nil {
					log.Println(err)
					http.Error(w, `{"reason": "invalid token"}`, http.StatusUnauthorized)
					return
				}
				if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
					userId := (*claims)["sub"].(string)
					ctx := context.WithValue(req.Context(), userIdKey{}, userId)
					next.ServeHTTP(w, req.WithContext(ctx))
				} else {
					next.ServeHTTP(w, req)
					return
				}
			} else {
				http.Error(w, `{"reason":"bearer is not found"}`, http.StatusUnauthorized)
				return
			}
		}

	})
}
func GetUserID(ctx context.Context) (string, bool) {
	switch v := ctx.Value(userIdKey{}).(type) {
	case string:
		return v, true
	default:
		return "", false
	}
}
