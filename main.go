package main

import (
	"fmt"
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/db"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/handlers"
	"github.com/gorilla/mux"
)

func init() {
	if err := config.ParseQuotaConfig("config/cofee_shop_quotas.yaml"); err != nil {
		fmt.Println(err)
	}
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "cofee_shop"
)

func main() {
	serverPort := ":8080"
	router := mux.NewRouter()
	sqlDB := db.SqlRepo{}
	if err := sqlDB.ConnectToDB(host, port, user, password, dbname); err != nil {
		panic(err)
	}

	cofeeHandlers := handlers.CofeeHandlers{Repo: &sqlDB}
	router.HandleFunc("/buycoffee", handlers.RequestValidation(cofeeHandlers.BuyCofee))
	fmt.Printf("server starts on localhost%s\n", serverPort)

	if err := http.ListenAndServe(serverPort, router); err != nil {
		fmt.Printf("server is down because: %s", err.Error())
	}
}
