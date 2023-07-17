package handlers

import (
	"fmt"
	"net/http"
)

func BuyCofee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only Get method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if !r.URL.Query().Has("cofeeType") {
		http.Error(w, "query parameter cofeeType is not specified", http.StatusUnprocessableEntity)
	}

	cofeeType := r.URL.Query().Get("cofeeType")

	fmt.Println(cofeeType)
}
