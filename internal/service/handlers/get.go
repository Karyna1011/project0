package handlers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/subgroup/project/resources"
	"net/http"
)

func GETHandler(w http.ResponseWriter, r *http.Request) {
	people, err := Person(r).Select()
	if err != nil {
		Log(r).WithError(err).Debug("error selecting persons")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response = make([]resources.PersonResponse, len(people))
	for i, p := range people {
		response[i] = p.Resource()
	}

	ape.Render(w, response)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
     message:= "Hello everybody"
	if err := json.NewEncoder(w).Encode(message);

	err != nil {
		panic(err)
	}
	ape.Render(w, message)
}



