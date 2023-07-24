package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/app"
)

var muMap sync.Map

func BuyCofee(w http.ResponseWriter, r *http.Request, bl *app.CofeeBL) {

	userId := r.Header.Get("user-id")
	mu, ok := muMap.Load(userId)
	if !ok {
		mu = &sync.Mutex{}
		muMap.Store(userId, &sync.Mutex{})
	}

	userMembership := r.Header.Get("membership-type")
	membership, err := bl.GetMembershipType(userMembership)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cofeeType := r.URL.Query().Get("cofeeType")
	if err := bl.ValidateCofeeType(cofeeType, membership); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.(*sync.Mutex).Lock()

	if err := bl.ProcessCofeeReq(r.Header.Get("user-id"), cofeeType, userMembership, membership.TypesOfCofee); err.Err != nil {
		if errors.Is(err.Err, app.ErrTooManyReq.Err) {
			errMsg := fmt.Sprintf("Hours to wait: %.2f", err.TimeToWait)
			http.Error(w, errMsg, err.StatusCode)
			mu.(*sync.Mutex).Unlock()
			return
		}

		http.Error(w, err.Err.Error(), http.StatusBadRequest)
	}

	mu.(*sync.Mutex).Unlock()
}
