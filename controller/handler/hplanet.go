package handler

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/carlosd-ss/star-wars/model"
	"github.com/carlosd-ss/star-wars/model/merrors"
	"github.com/carlosd-ss/star-wars/repo"
	"github.com/gorilla/mux"
)

func (rep *Repository) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var planet model.Planet

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro convertendo em JSON"))
		return
	}

	err = repo.Create(rep.Client, rep.Db, rep.Collection, planet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao Criar planeta"))
		return
	}
	json.NewEncoder(w).Encode(err)

}

func (rep *Repository) Getid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	res, err := repo.SearchForId(rep.Client, rep.Db, rep.Collection, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao buscar planeta"))
		return
	}
	json.NewEncoder(w).Encode(res)

}
func (rep *Repository) GetName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	name, _ := params["name"]
	res, err := repo.SearchForName(rep.Client, rep.Db, rep.Collection, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao buscar planeta pelo nome"))
		return
	}
	json.NewEncoder(w).Encode(res)

}
func (rep *Repository) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := repo.Delete(rep.Client, rep.Db, rep.Collection, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao deletar planeta"))
		return
	}

}
func (rep *Repository) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := repo.ListAll(rep.Client, rep.Db, rep.Collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao buscar planetas"))
		return
	}
	json.NewEncoder(w).Encode(res)

}
