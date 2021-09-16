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
	router.HandleFunc("/api/auth/login", login)
	router.HandleFunc("/api/user/register", registerUser)
	incomeRouter := router.PathPrefix("/api/income").Subrouter()
	incomeRouter.HandleFunc("/add", addIncome)
	incomeRouter.HandleFunc("/all", getAllIncomes)
	incomeRouter.HandleFunc("/latest", getLatestIncome)
	incomeRouter.Use(validateJWTMiddleware)
	categoryRouter := router.PathPrefix("/api/category").Subrouter()
	categoryRouter.HandleFunc("/add", addCategory)
	categoryRouter.HandleFunc("/all", getAllCategories)
	categoryRouter.Use(validateJWTMiddleware)
	fmt.Println("Listening port:8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
