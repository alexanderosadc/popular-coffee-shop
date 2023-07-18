package handlers

import (
	"net/http"
)

// RequestValidation middleware which validates query parameter cofeeType and
// headers user-id and membership-type.
func RequestValidation(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("cofeeType") {
			http.Error(w, "query parameter cofeeType is not specified", http.StatusUnprocessableEntity)
			return
		}

		userID := r.Header.Get("user-id")
		membershipType := r.Header.Get("membership-type")

		if len(userID) == 0 {
			http.Error(w, "missing user-id header", http.StatusBadRequest)
			return
		}

		if len(membershipType) == 0 {
			http.Error(w, "missing membership-type header", http.StatusBadRequest)
			return
		}

		f(w, r)
	}
}
