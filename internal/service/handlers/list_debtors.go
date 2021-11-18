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

func ListDebtors(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDebtorListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	debtors, err := Debtor(r).Select(req.OffsetPageParams)
	if err != nil {
		Log(r).WithError(err).Error("error selecting persons")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := resources.DebtorListResponse{
		Data: newDebtorsList(debtors),
	}

	ape.Render(w, response)
}

func newDebtorsList(debtors []data.Debtor) []resources.Debtor {
	result := make([]resources.Debtor, len(debtors))
	for i, debtor := range debtors {
		result[i] = newDebtorModel(debtor)
	}
	return result
}

func newDebtorModel(debtor data.Debtor) resources.Debtor {
	return resources.Debtor{
		Key: resources.Key{
			ID:   strconv.FormatInt(debtor.Id, 10),
			Type: resources.DEBTOR,
		},
		Attributes: resources.DebtorAttributes{
			Address: debtor.Address,
		},
	}
}
