package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
