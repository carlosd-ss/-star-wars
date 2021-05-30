package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/carlosd-ss/star-wars/model/merrors"

	"github.com/carlosd-ss/star-wars/model"
	"github.com/gorilla/mux"
)

func GetFilm(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, _ := params["id"]
	var film model.Film

	// consume api
	const url = "https://swapi.dev/api/films/"
	resp, err := http.Get(url + id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao buscar url"))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao ler planeta"))
		return
	}

	err = json.Unmarshal(body, &film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao converter planeta"))
		return
	}
	s, errm := json.Marshal(film)
	if errm != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao converter planeta"))
		return
	}
	_, errw := w.Write(s)
	if errw != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(merrors.FormatJSONError("Erro ao converter planeta"))
		return
	}
}
