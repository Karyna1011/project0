package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type PersonListRequest struct {
	pgdb.OffsetPageParams
}

func NewGetPersonListRequest(r *http.Request) (PersonListRequest, error) {
	request := PersonListRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
