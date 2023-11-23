package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1/todoservicev1connect"
	"github.com/Kimoto-Norihiro/connect-todo/server/database"
	"github.com/Kimoto-Norihiro/connect-todo/server/services/todoservice"
)

func main() {
	db, err := database.Conn()
	if err != nil {
		log.Println(err)
	}
	todoservice := todoservice.NewTODOService(db)
	mux := http.NewServeMux()
	mux.Handle(todoservicev1connect.NewTODOServiceHandler(todoservice))
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
