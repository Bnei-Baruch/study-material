package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/api/units", handleGetUnits).Methods("GET")
	router.HandleFunc("/api/unit", handleAddUnit).Methods("POST")

	router.Use(middleware)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
