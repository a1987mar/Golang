package middleware

import (
	"net/http"
)

type Middle func(http.Handler) http.Handler

// func Chain(middlewares ...Middleware) Middleware {
// 	return func(next http.Handler) http.Handler {
// 		for i := len(middlewares) - 1; i >= 0; i-- {
// 			next = middleware[i](next)
// 		}
// 		return next
// 	}
// }
