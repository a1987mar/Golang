package mark

import (
	"fmt"
	"strconv"

	"go/adv-demo/pkg/req"
	res "go/adv-demo/pkg/respose"
	"net/http"
)

type MarkHandler struct {
	MarkRepository *MarkRepository
}

type MarkHandlerDeps struct {
	MarkRepository *MarkRepository
}

func NewMarkHanler(router *http.ServeMux, deps MarkHandlerDeps) {
	handler := &MarkHandler{
		MarkRepository: deps.MarkRepository,
	}
	router.HandleFunc("GET /mark/{id}", handler.GetMark())
	router.HandleFunc("POST /mark", handler.CreateMark())
	router.HandleFunc("DELETE /mark/{id}", handler.DeleteMark())

}

func (h *MarkHandler) GetMark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		res.Json(w, id, http.StatusOK)
		fmt.Println(id)
	}
}

func (h *MarkHandler) CreateMark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[Mark](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		newmark := NewMark(body.Mark)
		res_, err := h.MarkRepository.Create(newmark)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, res_, http.StatusOK)
	}
}

func (h *MarkHandler) DeleteMark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idMark := r.PathValue("id")
		idMarkInt, err := strconv.Atoi(idMark)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resdel := h.MarkRepository.DeleteForId(idMarkInt)
		if resdel != nil {
			http.Error(w, resdel.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, resdel, http.StatusOK)

	}
}
