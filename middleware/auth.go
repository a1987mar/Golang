package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		auth := header.Get("Authorization")
		if auth == "" {
			fmt.Println("not filled")
			return
		}
		fmt.Println(auth)
		token := strings.TrimPrefix(auth, "Barier ")
		fmt.Println(token)
	})
}
