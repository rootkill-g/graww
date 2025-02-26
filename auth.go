package main

import (
	c "context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productId = r.URL.Query().Get("productId")
		productInfo, ok := database[productId]
		if !ok || productId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token := r.Header.Get("Authorization")
		if !isValidToken(productInfo, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := c.WithValue(r.Context(), "productInfo", productInfo)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

func isValidToken(productInfo ProductInfo, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == productInfo.Token
	}
	return false
}
