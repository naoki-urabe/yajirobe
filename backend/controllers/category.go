package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	if category.CategoryCode == "" || category.CategoryName == "" {
		w.WriteHeader(406)
		return
	}
	isSuccess := models.AddCategory(&category)
	if isSuccess == false {
		w.WriteHeader(406)
		return
	}
	responseBody, err := json.Marshal(category)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
})
var editCategory = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	editId, _ := vars["category"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var category models.Category
	if err := json.Unmarshal(reqBody, &category); err != nil {
		log.Fatal(err)
	}
	models.EditCategory(&category, editId)
	fmt.Fprintln(w, vars["category"])
})
var deleteCategory = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	deleteId, _ := vars["category"]
	models.DeleteCategory(deleteId)
	fmt.Fprintln(w, vars["category"])
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
