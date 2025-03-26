package link

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/middleware"
	"go/adv-demo/pkg/req"
	res "go/adv-demo/pkg/respose"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /link/{hash}", handler.Goto())
	router.HandleFunc("GET /link", handler.GetLink())
	router.HandleFunc("GET /link/getdomen/{domen}", handler.GetWhere())
}

func (c *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			panic(err.Error())
		}
		link := NewLink(body.Url)
		createdLink, err := c.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, 201)
	}
}
func (c *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}

		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			panic(err.Error())
		}
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := c.LinkRepository.UpdateByHash(&Link{
			Model: gorm.Model{
				ID: uint(idInt)},
			Url:  body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, 200)
	}
}
func (c *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		_, err = c.LinkRepository.GetById(idInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res_ := c.LinkRepository.DeleteById(idInt)
		if res_ != nil {
			http.Error(w, res_.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, res_, 200)
	}
}
func (c *LinkHandler) Goto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		getByHash, err := c.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, getByHash, 200)
	}
}

func (c *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ld_int, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
		offint, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		// res.Json(w, c.LinkRepository.GetALL(uint(ld_int), uint(offint)), 200)
		getLinks := GetAllLinkCount{
			Links:  c.LinkRepository.GetALL(uint(ld_int), uint(offint)),
			CountL: c.LinkRepository.Count(),
		}

		res.Json(w, getLinks, 201)
	}
}

func (c *LinkHandler) GetWhere() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domenid := r.PathValue("domen")

		res.Json(w, c.LinkRepository.GetWhereLink(domenid), 201)
	}
}
