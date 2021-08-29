package controllers

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"yajirobe/models"
)

var addIncome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var newIncome models.NewIncome
	if err := json.Unmarshal(reqBody, &newIncome); err != nil {
		log.Fatal(err)
	}
	models.AddIncome(&newIncome)
	responseBody, err := json.Marshal(newIncome)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
