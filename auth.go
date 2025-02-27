package main

import (
	c "context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			var productId = r.URL.Query().Get("productId")

			productInfo, ok := database[productId]

			if !ok || productId == "" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			tokenChecker(w, r)

			ctx := c.WithValue(r.Context(), "productInfo", productInfo)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		} else {
			tokenChecker(w, r)
			next.ServeHTTP(w, r)
		}
	}
}

func tokenChecker(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !isValidToken(token) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
}

func isValidToken(token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == adminTk.Token
	}
	return false
}
