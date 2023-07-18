package handlers

import (
	"fmt"
	"net/http"
)

func BuyCofee(w http.ResponseWriter, r *http.Request) {

	cofeeType := r.URL.Query().Get("cofeeType")

	fmt.Println(cofeeType)
}
