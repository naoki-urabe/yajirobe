package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"yajirobe/models"
)

func StartWebServer() error {
	models.ConnectDb()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/income/add", addIncome)
	router.HandleFunc("/api/category/add", addCategory)
	fmt.Println("Listening port:8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
