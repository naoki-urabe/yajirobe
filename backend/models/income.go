package models

import (
	_ "fmt"
	"log"
	"time"
)

var addIncomeQuery = `
INSERT INTO incomes (dt,summary,income,tag,user,month) VALUES(?,?,?,?,?,?);
`

var updateIncomeQuery = `
UPDATE incomes SET dt=?,summary=?,income=?,tag=?,month=? WHERE id = ?;
`

var deleteIncomeQuery = `
DELETE FROM incomes WHERE id = ?;`

var getAllIncomesQuery = `
SELECT * FROM incomes WHERE user = ?;`

var getLatestIncomeQuery = `
SELECT * FROM incomes WHERE user = ? ORDER BY id DESC LIMIT 1;
`

var getIncomeMonthsQuery = `
SELECT month FROM incomes GROUP BY month;
`

var getMonthlyIncomeQuery = `
SELECT * FROM incomes WHERE user = ? AND month = ?;
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

func UpdateIncome(income *Income, editId int) {
	_, err := Db.Queryx(updateIncomeQuery, income.Dt, income.Summary, income.Income, income.Tag, income.Month, editId)
	if err != nil {
		log.Println(err)
	}
}

func DeleteIncome(deleteId int) {
	_, err := Db.Queryx(deleteIncomeQuery, deleteId)
	if err != nil {
		log.Println(err)
	}
}

func GetAllIncomes(user string, incomes *[]Income) {
	Db.Select(incomes, getAllIncomesQuery, user)
}

func GetLatestIncome(user string, income *Income) {
	Db.Get(income, getLatestIncomeQuery, user)
}

func GetIncomeMonths(months *[]string) {
	Db.Select(months, getIncomeMonthsQuery)
}

func GetMonthlyIncomes(user string, month string, incomes *[]Income) {
	Db.Select(incomes, getMonthlyIncomeQuery, user, month)
}
