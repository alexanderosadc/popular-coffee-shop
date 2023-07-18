package handlers

import (
	"net/http"

	"github.com/alexanderosadc/popular-coffee-shop/config"
)

func BuyCofee(w http.ResponseWriter, r *http.Request) {

	membershipType := r.Header.Get("membership-type")

	quota, ok := config.Conf[membershipType]
	if !ok {
		http.Error(w, "there is no such membership", http.StatusBadRequest)
	}

	cofeeType := r.URL.Query().Get("cofeeType")
	_, ok = quota.TypesOfCofee[cofeeType]
	if !ok {
		http.Error(w, "there is no such type of cofee", http.StatusBadRequest)
	}
}
