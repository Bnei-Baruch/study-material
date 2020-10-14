package api

import (
	"database/sql"
	"github.com/Bnei-Baruch/study-material/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func handleAddUnit(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost/sm?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	u := &models.Unit{Title: "title 2", Description: "description 2"}
	err = u.Insert(context.Background(), db, boil.Infer())

	w.WriteHeader(http.StatusOK)
	w.Write(u.ID)
}
