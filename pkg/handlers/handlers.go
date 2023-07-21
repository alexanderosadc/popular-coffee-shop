package handlers

import (
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/app"
)

func BuyCofee(w http.ResponseWriter, r *http.Request, bl *app.CofeeBL) {
	userMembership := r.Header.Get("membership-type")
	membership, err := bl.GetMembershipType(userMembership)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cofeeType := r.URL.Query().Get("cofeeType")
	if err := bl.ValidateCofeeType(cofeeType, membership); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := bl.ProcessCofeeReq(r.Header.Get("user-id"), cofeeType, userMembership, membership.TypesOfCofee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
