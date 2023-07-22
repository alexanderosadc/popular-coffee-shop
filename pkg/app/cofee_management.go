package app

import (
	"errors"
	"time"

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

	selectedCofeeType, err := bl.getSelectedCofeeType(cofeeType, membershipCofees)
	if err != nil {
		return err
	}

	eligible, err := bl.isUserEligible(userID, selectedCofeeType)
	if err != nil {
		return err
	}

	if eligible {
		
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

// isUserEligible method is responsible for checking if user eligible for new cofee.
func (bl *CofeeBL) isUserEligible(id string, selectedCofeeType *domain.CofeeType) (bool, error) {
	purchasehistory, err := bl.Repo.GetPurchasesByUserID(id)
	if err != nil {
		return false, err
	}

	if len(purchasehistory) == 0 {
		return true, nil
	}

	nrOfAvialiablePurchase := 0
	for i := 0; i < selectedCofeeType.Limit; i++ {
		timeOfPurchase := purchasehistory[i].Time

		nextPurchaseTime, err := bl.calculateNextPurchaseTime(selectedCofeeType.TimeToRefresh, timeOfPurchase)
		if err != nil {
			return false, err
		}

		if !time.Now().Before(*nextPurchaseTime) {
			nrOfAvialiablePurchase++
		}
	}

	if nrOfAvialiablePurchase > selectedCofeeType.Limit {
		return false, errors.New("422")
	}

	return true, nil
}

// getSelectedCofeeType extracts cofee type object from config about based on cofeeType string
func (bl *CofeeBL) getSelectedCofeeType(cofeeType string, cofeeQuotas []domain.CofeeType) (*domain.CofeeType, error) {
	var selectedCofeeType *domain.CofeeType

	for _, cofee := range cofeeQuotas {
		if cofee.CofeeName == cofeeType {
			selectedCofeeType = &cofee
			break
		}
	}

	if selectedCofeeType == nil {
		return nil, errors.New("this cofee type not found in collection")
	}

	return selectedCofeeType, nil
}

// calculateNextPurchaseTime adds to the time of purchase delay when next purchase is avialiable
func (bl *CofeeBL) calculateNextPurchaseTime(timeToRefresh string, purchaseTime time.Time) (*time.Time, error) {
	durationToRefresh, err := time.ParseDuration(timeToRefresh)
	if err != nil {
		return nil, err
	}

	nextPurchaseTime := purchaseTime.Add(durationToRefresh)
	return &nextPurchaseTime, nil
}
