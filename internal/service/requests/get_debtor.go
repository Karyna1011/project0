package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type DebtorListRequest struct {
	pgdb.OffsetPageParams
}

func NewGetDebtorListRequest(r *http.Request) (DebtorListRequest, error) {
	request := DebtorListRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
