package middlewares

import (
	"context"
	"net/http"

	"github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/models"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := authHeader[len("Bearer "):]

		var user models.User
		err := db.DB.Where("token = ?", token).First(&user).Error
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(UserContextKey).(*models.User)
	return user, ok
}
