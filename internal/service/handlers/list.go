package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/subgroup/project/internal/data"
	"gitlab.com/tokend/subgroup/project/internal/service/requests"
	"gitlab.com/tokend/subgroup/project/resources"
	"net/http"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {

	req, err := requests.NewGetPersonListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	people, err := Person(r).Select(req.OffsetPageParams)
	if err != nil {
		Log(r).WithError(err).Error("error selecting persons")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := resources.PersonListResponse{
		Data: newPeopleList(people),
	}

	ape.Render(w, response)
}

func newPeopleList(people []data.Person) []resources.Person {
	result := make([]resources.Person, len(people))
	for i, person := range people {
		result[i] = newPersonModel(person)
	}
	return result
}

func newPersonModel(person data.Person) resources.Person {
	return resources.Person{
		Key: resources.Key{
			ID:   strconv.FormatInt(person.Id, 10),
			Type: resources.PERSON,
		},
		Attributes: resources.PersonAttributes{
			Name:      person.Name,
			Completed: person.Completed,
			Duration:  person.Duration,
		},
	}
}
