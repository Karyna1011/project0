package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/subgroup/project/resources"
)

type DebtorQ interface {
	New() DebtorQ
	Get() (*Debtor, error)
	Select(query pgdb.OffsetPageParams) ([]Debtor, error)
	Insert(data Debtor) (Debtor, error)
	Update(data Debtor) (Debtor, error)
	FilterById(data int64) DebtorQ
	FilterByAddress(data string) DebtorQ
	Delete(data string) error
}

type Debtor struct {
	Id      int64  `db:"id"        structs:"-"`
	Address string `db:"address"      structs:"address"`
}

func (p *Debtor) Resource() resources.DebtorResponse {
	return resources.DebtorResponse{
		Data: resources.Debtor{
			Attributes: resources.DebtorAttributes{
				Address: p.Address,
			},
			Key: resources.NewKeyInt64(p.Id, resources.DEBTOR),
		},
	}
}
