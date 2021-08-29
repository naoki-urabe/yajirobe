package models

import (
	"fmt"
	"time"
)

var addIncomeQuery = `
INSERT INTO incomes (dt,summary,income,tag) VALUES(?,?,?,?);
`

type Income struct {
	Id      int       `db:"id" json:"id"`
	Dt      time.Time `db:"dt" json:"dt"`
	Summary string    `db:"summary" json:"summary"`
	Income  int       `db:"income" json:"income"`
	Tag     string    `db:"tag" json:"tag"`
}

type NewIncome struct {
	Dt      time.Time `db:"dt" json:"dt"`
	Summary string    `db:"summary" json:"summary"`
	Income  int       `db:"income" json:"income"`
	Tag     string    `db:"tag" json:"tag"`
}

func AddIncome(newIncome *NewIncome) {
	Db.MustExec(addIncomeQuery, newIncome.Dt, newIncome.Summary, newIncome.Income, newIncome.Tag)
}
