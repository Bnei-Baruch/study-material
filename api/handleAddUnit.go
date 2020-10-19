package api

import (
	"encoding/json"
	"github.com/Bnei-Baruch/study-material/common"
	"github.com/Bnei-Baruch/study-material/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/net/context"
	"net/http"
)

func handleAddUnit(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("body")

	u := &models.Unit{Title: title, Description: description}
	err := u.Insert(context.Background(), common.DB, boil.Infer())
	common.FatalIfNil(err)

	res, errM := json.Marshal(common.PostResult{Id: u.UnitID})
	common.FatalIfNil(errM)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
