package main

import (
	"fmt"
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/handlers"
)

var port string = ":8080"

func init() {
	config.ParseQuotaConfig("config/cofee_shop_quotas.yaml")
}

func main() {
	http.HandleFunc("/buycoffee", handlers.RequestValidation(handlers.BuyCofee))
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("server is down because:" + err.Error())
	}
}
