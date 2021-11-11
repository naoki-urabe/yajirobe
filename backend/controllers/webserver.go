package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"yajirobe/config"
	"yajirobe/models"
)

func StartWebServer() error {
	models.ConnectDb()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/auth/login", login)
	router.HandleFunc("/api/user/register", registerUser)
	incomeRouter := router.PathPrefix("/api/income").Subrouter()
	incomeRouter.HandleFunc("/add", addIncome)
	incomeRouter.HandleFunc("/edit/{id}", editIncome)
	incomeRouter.HandleFunc("/delete/{id}", deleteIncome)
	incomeRouter.HandleFunc("/all", getAllIncomes)
	incomeRouter.HandleFunc("/latest", getLatestIncome)
	incomeRouter.HandleFunc("/months", getAllMonths)
	incomeRouter.HandleFunc("/monthly-incomes", getMonthlyIncome)
	incomeRouter.Use(validateJWTMiddleware)
	categoryRouter := router.PathPrefix("/api/category").Subrouter()
	categoryRouter.HandleFunc("/add", addCategory)
	categoryRouter.HandleFunc("/all", getAllCategories)
	categoryRouter.Use(validateJWTMiddleware)
	fmt.Printf("Listening port:%s...\n", config.Config.ApiPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.ApiPort), router)
}
