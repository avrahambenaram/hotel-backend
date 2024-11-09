package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
)

func ParseBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content type must be json", http.StatusForbidden)
			return
		}

		product := entity.Client{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ctx := context.WithValue(r.Context(), "client", product)
		*r = *r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
