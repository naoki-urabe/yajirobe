package models

import (
	_ "fmt"
	"time"
)

var addIncomeQuery = `
INSERT INTO incomes (dt,summary,income,tag,user,month) VALUES(?,?,?,?,?,?);
`

var getAllIncomesQuery = `
SELECT * FROM incomes WHERE user = ?;`

var getLatestIncomeQuery = `
SELECT * FROM incomes WHERE user = ? ORDER BY id DESC LIMIT 1;
`

type Income struct {
	Id      int       `db:"id" json:"id"`
	Dt      time.Time `db:"dt" json:"dt"`
	Summary string    `db:"summary" json:"summary"`
	Income  int       `db:"income" json:"income"`
	Tag     string    `db:"tag" json:"tag"`
	User    string    `db:"user" json:"user"`
	Month   string    `db:"month" json:"month"`
}

func AddIncome(income *Income) {
	Db.MustExec(addIncomeQuery, income.Dt, income.Summary, income.Income, income.Tag, income.User, income.Month)
}

func GetAllIncomes(user string, incomes *[]Income) {
	Db.Select(incomes, getAllIncomesQuery, user)
}

func GetLatestIncome(user string, income *Income) {
	Db.Get(income, getLatestIncomeQuery, user)
}
