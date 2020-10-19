package api

import (
	"context"
	"encoding/json"
	"github.com/Bnei-Baruch/study-material/common"
	"github.com/Bnei-Baruch/study-material/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"net/http"
)

func handleGetUnits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	u := &models.Unit{Title: "title 2", Description: "description 2"}
	err := u.Insert(context.Background(), common.DB, boil.Columns{})
	common.FatalIfNil(err)

	var (
		unitId      int
		title       string
		description string
	)

	rows, err := common.DB.Query("select unit_id, title, description from units")
	common.FatalIfNil(err)

	var units []common.UnitForClient
	for rows.Next() {
		err := rows.Scan(&unitId, &title, &description)
		common.FatalIfNil(err)

		units = append(units, common.UnitForClient{Title: title, Description: description})
		log.Println(unitId, title, description)
	}
	err = rows.Err()
	common.FatalIfNil(err)

	b, err := json.Marshal(units)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
