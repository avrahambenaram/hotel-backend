package middleware

import (
	"context"
	"encoding/json"
	"net/http"
)

func ParseBody[T any](next http.Handler, v T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content type must be json", http.StatusForbidden)
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&v); err != nil {
			http.Error(w, "Erro no json", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ctx := context.WithValue(r.Context(), "body", v)
		*r = *r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
