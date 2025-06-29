package main

import (
	"google-login/internal/handler/rest"
	"google-login/internal/repository"
	"google-login/internal/service"
	"google-login/pkg/bcrypt"
	"google-login/pkg/config"
	"google-login/pkg/database/mariadb"
	"google-login/pkg/jwt"
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

	r := rest.NewRest(svc)
	r.MountEndpoint()
	r.Run()
}
