package controllers

import (
	"log"
	"yajirobe/models"
	"net/http"
	"encoding/json"
	"io/ioutil"
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
