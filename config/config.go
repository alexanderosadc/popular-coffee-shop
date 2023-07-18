package config

import (
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

func ParseQuotaConfig(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(dat, &Conf); err != nil {
		return err
	}

	return nil
}
