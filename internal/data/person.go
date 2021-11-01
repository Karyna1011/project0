package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/subgroup/project/resources"
)

type PersonQ interface {
	New() PersonQ
	Get() (*Person, error)
	//Select() ([]Person, error)
	Select(query pgdb.OffsetPageParams) ([]Person, error)
	Insert(data Person) (Person, error)
	Update(data Person) (Person, error)
	FilterById(data int64) PersonQ
}

type Person struct {
	Id        int64  `db:"id"        structs:"-"`
	Name      string `db:"name"      structs:"name"`
	Completed bool   `db:"completed" structs:"completed"`
	Duration  int64  `db:"duration"  structs:"duration"`
}

func (p *Person) Resource() resources.PersonResponse {
	return resources.PersonResponse{
		Data: resources.Person{
			Attributes: resources.PersonAttributes{
				Name:      p.Name,
				Completed: p.Completed,
				Duration:  p.Duration,
			},
			Key: resources.NewKeyInt64(p.Id, resources.PERSON),
		},
	}
}
