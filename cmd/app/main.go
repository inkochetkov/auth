package main

import (
	"log"
	"net/http"

	"github.com/inkochetkov/auth/internal/app/base"
	"github.com/inkochetkov/auth/internal/entity"
	"github.com/inkochetkov/auth/internal/external/sqlite"
	serv "github.com/inkochetkov/auth/internal/server"
	"github.com/inkochetkov/config/pkg/config"
)

func main() {

	c, err := config.NewConfig("config/config.yaml", &entity.Config{})
	if err != nil {
		log.Fatal("fail config", err)
	}
	conf := c.(*entity.Config)

	app, clean, err := start(*conf)
	if err != nil {
		log.Fatal("fail start serv", err)
	}
	defer clean()

	srv := &http.Server{
		Addr:         conf.HTTP.Port,
		Handler:      app.http,
		ReadTimeout:  conf.HTTP.Timeout.Read,
		WriteTimeout: conf.HTTP.Timeout.Write,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("fail http server", err)
	}

}

type application struct {
	http http.Handler
}

func new(http http.Handler) application {
	return application{http: http}
}

func start(config entity.Config) (application, func(), error) {
	sqlite := sqlite.New(config)
	api := base.NewAPI(config, sqlite)
	router := serv.NewRouter(api, config)
	handler := serv.NewHTTP(config, router)

	mainApplication := new(handler)
	return mainApplication, func() {}, nil
}
