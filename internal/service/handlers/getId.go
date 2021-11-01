package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
)

func GetByIndex(w http.ResponseWriter, r *http.Request){
	IdString := chi.URLParam(r, "id")

	personQ := Person(r)

	Id, err := strconv.ParseInt(IdString, 10, 64)

	if err != nil {
		Log(r).WithError(err).Error("failed to parse mission id")
		ape.Render(w, problems.InternalError())
		return
	}

	person, err := personQ.FilterById(Id).Get()

	if err != nil {
		Log(r).WithError(err).Error("failed to get person")
		ape.Render(w, problems.InternalError())
		return
	}

	if person == nil{
		ape.Render(w, problems.NotFound())
		return
	}

	result:=person

	ape.Render(w, result)

}