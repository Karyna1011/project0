package postgres

import (
	"database/sql"
	"github.com/fatih/structs"
	"gitlab.com/tokend/subgroup/project/internal/data"

	"github.com/Masterminds/squirrel"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const tablePerson = "person"

type personQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewPersonQ(db *pgdb.DB) data.PersonQ {
	return &personQ{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(tablePerson),
	}
}

func (d *personQ) New() data.PersonQ {
	return NewPersonQ(d.db)
}

func (d *personQ) Get() (*data.Person, error) {
	var result data.Person

	err := d.db.Get(&result, d.sql)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (d *personQ) Select(query pgdb.OffsetPageParams) ([]data.Person, error) {
	var result []data.Person

	err := d.db.Select(&result, query.ApplyTo(d.sql, "id"))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func (d *personQ) Insert(person data.Person) (data.Person, error) {
	clauses := structs.Map(person)

	query := squirrel.Insert(tablePerson).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&person, query)
	if err != nil {
		return data.Person{}, err
	}

	return person, err

}

func (d *personQ) Update(person data.Person) (data.Person, error) {
	clauses := structs.Map(person)

	query := squirrel.Update(tablePerson).Where(squirrel.Eq{"id": person.Id}).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&person, query)
	if err != nil {
		return data.Person{}, err
	}

	return person, err
}

func (d personQ) FilterById(Id int64) data.PersonQ {
	d.sql = d.sql.Where(squirrel.Eq{"id": Id})

	return &d
}

func (d personQ) FilterByAddress(Address string) data.PersonQ {
	d.sql = d.sql.Where(squirrel.Eq{"address": Address})

	return &d
}

/*func (q *personQ) GetNotActiveAddressesAmount() (int64, error) {
	var result int64
	stmt := q.sql.Select("COUNT (*)").From(addressTableName).Where(sq.Eq{"active": false})
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return result, nil
}*/

/*func (d personQ) MaxId(Address string) data.PersonQ {
	//f:=squirrel.Select(*).Where()

	d.sql = d.sql.Where(squirrel.Eq{"id": })

	return &d
}

func (d *personQ) GetAmount() (int64, error) {
	var result int64
	stmt := squirrel.Select("*").From(tablePerson)
	err := d.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return result, nil
}*/
