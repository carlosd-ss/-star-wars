package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	mcon "github.com/carlosd-ss/-star-wars/pkg/mongo"

	"github.com/carlosd-ss/-star-wars/controller/route"

	cfg "github.com/carlosd-ss/-star-wars/pkg/config"
	"github.com/carlosd-ss/-star-wars/repo"
)

func main() {
	s := repo.Service{
		R: &repo.Repository{
			Client:     mcon.Connect(),
			Db:         cfg.Db,
			Collection: cfg.Collection,
		},
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		if err := s.R.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      route.Router(),
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
