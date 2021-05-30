package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/carlosd-ss/star-wars/controller/handler"

	"github.com/carlosd-ss/star-wars/controller/route"

	cfg "github.com/carlosd-ss/star-wars/pkg/config"
)

func main() {
	s := handler.Service{
		R: &handler.Repository{
			Client:     cfg.Connect,
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
		Addr:         cfg.Port,
		Handler:      route.Router(s.R),
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
