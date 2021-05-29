package handler

import (
	"encoding/json"
	"net/http"

	"github.com/carlosd-ss/-star-wars/model"
	"github.com/carlosd-ss/-star-wars/model/merrors"
	"github.com/carlosd-ss/-star-wars/repo"
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var repository *repo.Repository
	var planet model.Planet

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro convertendo em JSON"))
		return
	}

	err = repository.Create(planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(err)
	w.WriteHeader(200)

}

//func Getid(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var repository repo.Repository
//	var params = mux.Vars(r)
//
//	id, _ := primitive.ObjectIDFromHex(params["id"])
//	res, err := repository.SearchForId(id)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write(merrors.FormatJSONError(err.Error()))
//		return
//	}
//	json.NewEncoder(w).Encode(res)
//	w.WriteHeader(200)
//
//}
