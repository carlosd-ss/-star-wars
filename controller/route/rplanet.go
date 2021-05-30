package route

import (
	"github.com/carlosd-ss/star-wars/controller/handler"
	"github.com/gorilla/mux"
)

func Router(rep *handler.Repository) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/v1/planet", rep.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/planet/{id}", rep.Getid).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/planet/search/{name}", rep.GetName).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/planet/films/{id}", handler.GetFilm).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/planet", rep.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/planet/{id}", rep.Delete).Methods("DELETE", "OPTIONS")

	return router
}
