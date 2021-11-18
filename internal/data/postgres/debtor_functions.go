package postgres

import (
	"database/sql"
	"github.com/fatih/structs"
	"gitlab.com/tokend/subgroup/project/internal/data"

	"github.com/Masterminds/squirrel"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const tableDebtor = "debtor"

type debtorQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewDebtorQ(db *pgdb.DB) data.DebtorQ {
	return &debtorQ{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(tableDebtor),
	}
}

func (d *debtorQ) New() data.DebtorQ {
	return NewDebtorQ(d.db)
}

func (d *debtorQ) Get() (*data.Debtor, error) {
	var result data.Debtor

	err := d.db.Get(&result, d.sql)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (d *debtorQ) Select(query pgdb.OffsetPageParams) ([]data.Debtor, error) {
	var result []data.Debtor

	err := d.db.Select(&result, query.ApplyTo(d.sql, "id"))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func (d *debtorQ) Insert(debtor data.Debtor) (data.Debtor, error) {
	clauses := structs.Map(debtor)

	query := squirrel.Insert(tableDebtor).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&debtor, query)
	if err != nil {
		return data.Debtor{}, err
	}

	return debtor, err

}

func (d *debtorQ) Update(debtor data.Debtor) (data.Debtor, error) {
	clauses := structs.Map(debtor)

	query := squirrel.Update(tableDebtor).Where(squirrel.Eq{"id": debtor.Id}).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&debtor, query)
	if err != nil {
		return data.Debtor{}, err
	}

	return debtor, err
}

func (d debtorQ) FilterById(Id int64) data.DebtorQ {
	d.sql = d.sql.Where(squirrel.Eq{"id": Id})

	return &d
}

func (d debtorQ) FilterByAddress(Address string) data.DebtorQ {
	d.sql = d.sql.Where(squirrel.Eq{"address": Address})

	return &d
}

func (d *debtorQ) Delete(Address string) error {
	query := squirrel.Delete(tableDebtor).Where(squirrel.Eq{"address": Address})
	err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
