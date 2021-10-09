package controllers

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"yajirobe/models"
)

type User struct {
	User string
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
