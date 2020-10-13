package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleAddUnit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Collection id: %v\n", vars["cId"])
}
