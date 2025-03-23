package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	Statuscode int
}

func (w *WrapperWriter) WriteHeader(Statuscode int) {
	w.ResponseWriter.WriteHeader(Statuscode)
	w.Statuscode = Statuscode
}
