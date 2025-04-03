package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/mark"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"net/http"
)

func main() {
	//
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	//Repositories
	linkrep := link.NewLinkRepository(db)
	regrep := auth.NewRegRepository(db)
	markrep := mark.NewMarkRepository(db)
	userRep := user.NewUserRepository(db)
	newauth := auth.NewAuthService(userRep)
	//-->Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:        conf,
		AuthService:   newauth,
		RegRepository: regrep,
	})
	link.NewAuthHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkrep,
		Config:         conf,
	})
	mark.NewMarkHanler(router, mark.MarkHandlerDeps{
		MarkRepository: markrep,
	})
	user.NewUserHandler(router, user.UserHandlerDepo{
		UserRepos: userRep,
	})

	//<--Handler

	server := http.Server{
		Addr: conf.DB.Port,
		//Handler: middleware.IsAuthed(middleware.Logging(router)),
		Handler: router,
	}
	fmt.Println("server start port:8081")
	server.ListenAndServe()
}
