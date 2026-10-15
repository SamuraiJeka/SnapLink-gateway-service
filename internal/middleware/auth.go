package middleware

import (
	"context"
	"net/http"
	"strings"

	keyprovider "github.com/SamuraiJeka/SnapLink-gateway-service/internal/key_provider"
)

type AuthMiddleware struct {
	verifier *keyprovider.Verifier
}

func NewAuthMiddleware(
	verifier *keyprovider.Verifier,
) *AuthMiddleware {
	return &AuthMiddleware{
		verifier: verifier,
	}
}

const UserIdKey = "user_id"

func (m *AuthMiddleware) Handle(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "authorization required", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		
		if !strings.HasPrefix(header, prefix) {
			http.Error(w, "invalid autorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, prefix)

		userId, err := m.verifier.Verify(r.Context(), token)
		if err != nil {
			http.Error(w, "invalid access token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserIdKey,
			userId,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
