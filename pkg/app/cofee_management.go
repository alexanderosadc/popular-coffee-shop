package app

import (
	"errors"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/db"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
)

type CofeeBL struct {
	Repo db.Repository
}

// GetMembershipType compares existing membership types from config to membership introduced by user
// returns membershipType object if membershipType exists
func (bl *CofeeBL) GetMembershipType(membershipType string) (*domain.MembershipType, error) {
	membershipQuota, ok := config.Conf[membershipType]
	if !ok {
		return nil, errors.New("there is no such membership")
	}

	return &membershipQuota, nil
}

// ValidateCofeeType verifies if the cofeeType introduced by user exists in config
func (bl *CofeeBL) ValidateCofeeType(cofeeType string, membership *domain.MembershipType) error {
	for _, item := range membership.TypesOfCofee {
		if item.CofeeName == cofeeType {
			return nil
		}
	}

	return errors.New("there is no such type of cofee")
}

func (bl *CofeeBL) GetCofee(userID string, cofeeType string, membership *domain.MembershipType) error {
	return nil
}
