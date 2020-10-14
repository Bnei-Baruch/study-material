package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Bnei-Baruch/study-material/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"net/http"
)

type UnitForClient struct {
	title       string `json: "title"`
	description string `json: "description"`
}

func handleGetUnits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)


	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost/sm?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	u := &models.Unit{Title: "title 2", Description: "description 2"}
	err = u.Insert(context.Background(), db,boil.Infer())



	var (
		id          int
		title       string
		description string
	)

	rows, err := db.Query("select id, title, description from units")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	units := []UnitForClient{}
	for rows.Next() {
		err := rows.Scan(&id, &description, &title)
		if err != nil {
			log.Fatal(err)
		}

		units = append(units, UnitForClient{title: title, description: description})
		log.Println(id, title, description)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(units)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}
