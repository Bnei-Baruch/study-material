package api

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
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
