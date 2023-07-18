package main

import (
	"fmt"
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/handlers"
	"github.com/gorilla/mux"
)

func init() {
	if err := config.ParseQuotaConfig("config/cofee_shop_quotas.yaml"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/buycoffee", handlers.RequestValidation(handlers.BuyCofee))
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("server is down because: %s", err.Error())
	}
}
