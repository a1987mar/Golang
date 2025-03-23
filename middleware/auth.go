package middleware

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		auth := header.Get("Authorization")
		if auth == "" {
			fmt.Println("not filled")
			return
		}
		fmt.Println(auth)
		token := strings.TrimPrefix(auth, "Barier ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).ParseJWT(token)
		fmt.Println(isValid)
		fmt.Println(data)
		fmt.Println(token)
	})
}
