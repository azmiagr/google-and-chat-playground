package main

import (
	"google-login/internal/handler/rest"
	"google-login/internal/handler/websocket"
	"google-login/internal/repository"
	"google-login/internal/service"
	"google-login/pkg/bcrypt"
	"google-login/pkg/config"
	"google-login/pkg/database/mariadb"
	"google-login/pkg/jwt"
	"google-login/pkg/middleware"
	"log"
)

func main() {
	config.LoadEnvironment()
	oauth := config.NewOAuthConfig()
	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()

	svc := service.NewService(repo, bcrypt, jwt, oauth)

	hub := websocket.NewHub()
	go hub.Run()

	wsHandler := websocket.NewWebSocketHandler(hub)
	middleware := middleware.Init(svc, jwt)

	r := rest.NewRest(svc, wsHandler, middleware)
	r.MountEndpoint()
	r.Run()
}
