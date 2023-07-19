package config

import (
	"os"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
	"gopkg.in/yaml.v3"
)

var Conf map[string]domain.MembershipType

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
