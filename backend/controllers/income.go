package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"yajirobe/models"
)

type User struct {
	User string
}

func calcMonth(date time.Time) string {
	startYear := date.Year()
	endYear := date.Year()
	m := int(date.Month())
	lastMonth := m - 1
	if lastMonth == 0 {
		startYear -= 1
		lastMonth = 12
	}
	startDay, _ := time.Parse("2006-1-2", fmt.Sprintf("%d-%d-%d", startYear, lastMonth, 27))
	endDay, _ := time.Parse("2006-1-2", fmt.Sprintf("%d-%d-%d", endYear, m, 27))
	month := ""
	if startDay.Before(date) && date.Before(endDay) {
		month = fmt.Sprintf("%d-%02d", endYear, m)
	} else {
		if m+1 > 12 {
			endYear += 1
			m = 0
		}
		month = fmt.Sprintf("%d-%02d", endYear, m+1)
	}
	return month
}

var addIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var income models.Income
	if err := json.Unmarshal(reqBody, &income); err != nil {
		log.Fatal(err)
	}
	income.Month = calcMonth(income.Dt)
	models.AddIncome(&income)
	responseBody, err := json.Marshal(income)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
var editIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	editId, _ := strconv.Atoi(vars["id"])
    reqBody, _ := ioutil.ReadAll(r.Body)
	var income models.Income
	if err := json.Unmarshal(reqBody, &income); err != nil {
		panic(err)
	}
	income.Month = calcMonth(income.Dt)
	models.UpdateIncome(&income, editId)
	fmt.Fprintln(w, vars["id"])
})
var deleteIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deleteId, _ := strconv.Atoi(vars["id"])
	models.DeleteIncome(deleteId)
	fmt.Fprintln(w, vars["id"])
})

var getAllIncomes = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		panic(err)
	}
	var incomes []models.Income
	models.GetAllIncomes(user.User, &incomes)
	responseBody, err := json.Marshal(incomes)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getLatestIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		panic(err)
	}
	var income models.Income
	models.GetLatestIncome(user.User, &income)
	responseBody, err := json.Marshal(income)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)

})

var getAllMonths = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var months []string
	models.GetIncomeMonths(&months)
	responseBody, err := json.Marshal(months)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getMonthlyIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	type Cols struct {
		User  string
		Month string
	}
	var c Cols
	if err := json.Unmarshal(reqBody, &c); err != nil {
		panic(err)
	}
	var incomes []models.Income
	models.GetMonthlyIncomes(c.User, c.Month, &incomes)
	responseBody, err := json.Marshal(incomes)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
