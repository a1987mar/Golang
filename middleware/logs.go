package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			Statuscode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.Println(r.Method, r.URL.Path, t, wrapper.Statuscode)
	})
}
