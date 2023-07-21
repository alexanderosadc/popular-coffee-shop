package app

import (
	"errors"
	"fmt"

	"github.com/alexanderosadc/popular-coffee-shop/config"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/db"
	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
	"gorm.io/gorm"
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

// ProcessCofeeReq is method for
func (bl *CofeeBL) ProcessCofeeReq(userID, cofeeType, membership string, membershipCofees []domain.CofeeType) error {
	_, err := bl.getUser(userID, membership)
	if err != nil {
		return err
	}

	_, err = bl.checkQuota(userID, cofeeType, membershipCofees)
	if err != nil {
		return err
	}

	return nil
}

// getUser method is responsible for finding or creating user in DB. User is returned.
func (bl *CofeeBL) getUser(userID, membership string) (*domain.User, error) {
	user, err := bl.Repo.GetByID(userID)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user != nil {
		user.Membership = membership
		return user, nil
	}

	user = &domain.User{ID: userID, Membership: membership}

	if err = bl.Repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// checkQuota method is responsible for checking if user eligible for new cofee
func (bl *CofeeBL) checkQuota(userID, cofeeType string, membershipCofees []domain.CofeeType) (bool, error) {
	var selectedCofeeType *domain.CofeeType

	for _, cofee := range membershipCofees {
		if cofee.CofeeName == cofeeType {
			selectedCofeeType = &cofee
			break
		}
	}

	if selectedCofeeType == nil {
		return false, errors.New("this cofee type not found in collection")
	}

	history, err := bl.Repo.GetPurchasesByUserID(userID)
	if err != nil {
		return false, err
	}

	fmt.Println(history)
	return false, nil
}
