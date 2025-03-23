package middleware

import (
	"context"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

type keyEmail string

const (
	ContextEmailKey keyEmail = "ContextEmailKey"
)

func writeUnathed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Barier ") {
			writeUnathed(w)
			return
		}
		token := strings.TrimPrefix(auth, "Barier ")
		isvalid, data := jwt.NewJWT(config.Auth.Secret).ParseJWT(token)
		if !isvalid {
			writeUnathed(w)
			return
		}

		r.Context()
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
