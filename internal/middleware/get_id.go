package middleware

import (
	"context"
	"net/http"
	"strconv"
)

func GetId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathID := r.PathValue("ID")
		id, err := strconv.Atoi(pathID)
		if err != nil {
			http.Error(w, "Insert a valid ID (non negative integer)", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "id", uint(id))
		*r = *r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
