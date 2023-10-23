// /Users/spulakan/CloudProject/webapplication_go/api/middleware/auth_middleware.go

package middleware

import (
	"context"
	"net/http"

	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

type AuthMiddlewareDependencies struct {
	AuthHandlerDependencies handlers.AuthHandlerDependencies
}

func BasicAuthMiddleware(deps AuthMiddlewareDependencies, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Authenticate the user based on the provided username and password
		user, err := handlers.AuthenticateUser(deps.AuthHandlerDependencies, email, password)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication is successful, you can proceed to the next handler
		// and attach the authenticated user to the request context.
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
