package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/dto"
// 	"github.com/SamuraiJeka/SnapLink-gateway-service/internal/service"
// )

// type AuthMiddleware struct {
// 	service service.AuthService
// }

// func NewAuthMiddleware(
// 	service service.AuthService,
// ) *AuthMiddleware {
// 	return &AuthMiddleware{
// 		service: service,
// 	}
// }

// func (m *AuthMiddleware) Handle(
// 	next http.Handler,
// ) http.Handler {
// 	return http.HandlerFunc(func(
// 		w http.ResponseWriter,
// 		r *http.Request,
// 	) {
// 		header := r.Header.Get("Authorization")
// 		if header == "" {
// 			http.Error(w, "authorization required", http.StatusUnauthorized)
// 			return
// 		}

// 		const prefix = "Bearer "
		
// 		if !strings.HasPrefix(header, prefix) {
// 			http.Error(w, "invalid autorization header", http.StatusUnauthorized)
// 			return
// 		}

// 		token := strings.TrimPrefix(header, prefix)

// 		resp, err := m.service.ValidateAccessToken(
// 			r.Context(),
// 			dto.ValidateRequest{
// 				Access_token: token,
// 			},
// 		)
// 		if err != nil || !resp.Valid {
// 			http.Error(w, "invalid access token", http.StatusUnauthorized)
// 			return
// 		}

// 		ctx := setUserId(
// 			r.Context(),
// 			resp.User_id,
// 		)

// 		next.ServeHTTP(
// 			w,
// 			r.WithContext(ctx),
// 		)
// 	})
// }
