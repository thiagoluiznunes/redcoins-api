package internal

import (
	"context"
	"net/http"
	"strings"
)

// AuthorizeMiddleware : middleware to filter unauthorized requests
func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer ErrorsHandler(w, r)

		header := r.Header.Get("Authorization")
		if header == "" {
			ResponseHandler(w, r, 401, "Unauthorized access.")
			return
		}

		token := strings.Split(header, "Bearer")[1]
		isValid, signature := ValidateToken(strings.TrimSpace(token))

		if isValid {
			ctx := context.WithValue(r.Context(), "signature", signature)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		ResponseHandler(w, r, 401, "Unauthorized access.")
		return
	})
}
