package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/subgroup/project/resources"
)

type PersonQ interface {
	New() PersonQ
	Get() (*Person, error)
	Select(query pgdb.OffsetPageParams) ([]Person, error)
	Insert(data Person) (Person, error)
	Update(data Person) (Person, error)
	FilterById(data int64) PersonQ
	FilterByAddress(data string) PersonQ
	//GetAmount() (int64, error)

}

type Person struct {
	Id      int64  `db:"id"        structs:"-"`
	Address string `db:"address"      structs:"address"`
}

func (p *Person) Resource() resources.PersonResponse {
	return resources.PersonResponse{
		Data: resources.Person{
			Attributes: resources.PersonAttributes{
				Address: p.Address,
			},
			Key: resources.NewKeyInt64(p.Id, resources.PERSON),
		},
	}
}
