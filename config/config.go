package config

import (
	"fmt"
	"os"

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

var Conf []Quota

func ParseQuotaConfig(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := yaml.Unmarshal(dat, &Conf); err != nil {
		fmt.Println(err)
		return
	}
}
