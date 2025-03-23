package req

import (
	res "go/adv-demo/pkg/respose"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := DecodeBody[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 401)
		return nil, err
	}

	err = validBody(body)
	if err != nil {
		res.Json(*w, err.Error(), 401)
		return nil, err
	}

	return &body, nil
}
