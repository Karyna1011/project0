package handlers

import (
	"encoding/json"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"gitlab.com/tokend/subgroup/project/resources"
	"net/http"
)

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	var p resources.PersonResponse
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = Person(r).Insert(data.Person{
		Name:      p.Data.Attributes.Name,
		Completed: p.Data.Attributes.Completed,
		Duration:  p.Data.Attributes.Duration,
	})

	if err != nil {
		Log(r).WithError(err).Debug("error inserting person")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

