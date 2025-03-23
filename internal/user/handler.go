package user

import (
	"go/adv-demo/pkg/req"
	res "go/adv-demo/pkg/respose"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserRepos *UserRepository
}

type UserHandlerDepo struct {
	UserRepos *UserRepository
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDepo) {
	handler := &UserHandler{
		UserRepos: deps.UserRepos,
	}
	router.HandleFunc("POST /user", handler.Create())
	router.HandleFunc("GET  /user", handler.FindBy())
	router.HandleFunc("DELETE /user/{id}", handler.DeleteID())
}

func (u *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		usN := NewUser(body.Email, body.Name, body.Password)
		result, err := u.UserRepos.CreateUser(usN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, result, http.StatusOK)
	}
}
func (u *UserHandler) CreateRegister(email, password, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		us := NewUser(email, name, password)
		resul, err := u.UserRepos.CreateUser(us)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, resul, http.StatusOK)
	}
}

func (u *UserHandler) FindBy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[User_](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := u.UserRepos.FindByEmail(body.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, result, http.StatusOK)
	}
}

func (u *UserHandler) DeleteID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idDel := u.UserRepos.DeleteByID(idInt)
		if idDel != nil {
			http.Error(w, idDel.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, idDel, http.StatusOK)
	}
}
