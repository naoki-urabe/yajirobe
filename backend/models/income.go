package models

import (
	_ "fmt"
	"time"
)

var addIncomeQuery = `
INSERT INTO incomes (dt,summary,income,tag) VALUES(?,?,?,?);
`

var getAllIncomesQuery = `
SELECT * FROM incomes`

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
	// addIncome =
	Db.MustExec(addIncomeQuery, newIncome.Dt, newIncome.Summary, newIncome.Income, newIncome.Tag)
	/*id, _ := addIncome.LastInsertId()
	return int(id)*/
}

func GetAllIncomes(incomes *[]Income) {
	Db.Select(incomes, getAllIncomesQuery)
}

/*func GetLatestIncome(income []Income) {

}*/
