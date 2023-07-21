package app

import (
	"errors"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
)

// GetMembershipType compares existing membership types from config to membership introduced by user
// returns membershipType object if membershipType exists
func GetMembershipType(membershipType string) (*domain.MembershipType, error) {
	membershipQuota, ok := config.Conf[membershipType]
	if !ok {
		return nil, errors.New("there is no such membership")
	}

	return &membershipQuota, nil
}

// ValidateCofeeType verifies if the cofeeType introduced by user exists in config
func ValidateCofeeType(cofeeType string, membership *domain.MembershipType) error {
	for _, item := range membership.TypesOfCofee {
		if item.CofeeName == cofeeType {
			return nil
		}
	}

	return errors.New("there is no such type of cofee")
}

func ProcessCofeeRequest(userID string, cofeeType string, membership *domain.MembershipType) error {
	return nil
}
