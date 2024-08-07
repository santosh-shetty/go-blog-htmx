package middleware

import (
	"net/http"

	"github.com/santosh-shetty/blog/pkg/helpers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// Redirect to login page
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		tokenString := cookie.Value
		_, err = helpers.VerifyJWTToken(tokenString)
		if err != nil {
			// Redirect to login page
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// If token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
