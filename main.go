package main

import (
	"fmt"
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/handlers"
)

// init parses config files before start of the server
func init() {

}

var port string = ":8080"

func main() {
	http.HandleFunc("/buycoffee", handlers.BuyCofee)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("server is down because:" + err.Error())
	}
}
