package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"yajirobe/models"
)

var addCategory = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var category models.Category
	if err := json.Unmarshal(reqBody, &category); err != nil {
		log.Fatal(err)
	}
	models.AddCategory(&category)
	responseBody, err := json.Marshal(category)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})

var getAllCategories = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var categories []models.Category
	models.GetAllCategories(&categories)
	responseBody, err := json.Marshal(categories)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
