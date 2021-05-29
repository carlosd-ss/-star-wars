package route

import (
	"github.com/carlosd-ss/-star-wars/controller/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/v1/planet", handler.Post).Methods("POST", "OPTIONS")
	//router.HandleFunc("/v1/planet/{id}", handler.Getid).Methods("GET", "OPTIONS")
	//router.HandleFunc("/v1/planet", handler.GetAllUser).Methods("GET", "OPTIONS")
	//router.HandleFunc("/v1/planet/{id}", handler.UpdateUser).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/v1/planet/{id}", handler.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
