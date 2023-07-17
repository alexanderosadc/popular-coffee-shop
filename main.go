package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/handlers"
	"gopkg.in/yaml.v3"
)

type CofeeLimit struct {
	Time  string `yaml:"time"`
	Limit int    `yaml:"limit"`
}
type Quota struct {
	QuotaName    string                `yaml:"quota_name"`
	TypesOfCofee map[string]CofeeLimit `yaml:"types_of_cofee"`
}

// init parses config files before start of the server
func init() {
	dat, err := os.ReadFile("config/cofee_shop_quotas.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	yamlConf := []Quota{}

	if err := yaml.Unmarshal(dat, &yamlConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(yamlConf)
}

var port string = ":8080"

func main() {
	http.HandleFunc("/buycoffee", handlers.BuyCofee)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("server is down because:" + err.Error())
	}
}
