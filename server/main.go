package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1/todoservicev1connect"
	"github.com/Kimoto-Norihiro/connect-todo/server/database"
	"github.com/Kimoto-Norihiro/connect-todo/server/services/todoservice"

	"github.com/rs/cors"
)

func main() {
	db, err := database.Conn()
	if err != nil {
		log.Println(err)
	}
	todoservice := todoservice.NewTODOService(db)
	mux := http.NewServeMux()

	path, handler := todoservicev1connect.NewTODOServiceHandler(todoservice)
	mux.Handle(path, handler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Connect-Protocol-Version", "Content-Type"},
	})
	
	corsHandler := c.Handler(h2c.NewHandler(mux, &http2.Server{}))
	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
}
