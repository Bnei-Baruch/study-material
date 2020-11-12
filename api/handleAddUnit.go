package api

import (
	"encoding/json"
	cm "github.com/Bnei-Baruch/study-material/common"
	m "github.com/Bnei-Baruch/study-material/models"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

func handleAddUnit(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print("can't create unit wrong request body", err)
	}

	var last, err = m.Units(OrderBy(`created_at desc`)).One(context.Background(), cm.DB)
	if err != nil {
		log.Print("problem with connect to DB", err)
	}
	if last != nil {
		log.Printf("Last saved on DB: %d, %s, %s", last.UnitID, last.CreatedAt, last.Title)
	}

	createdAt, err := time.Parse(viper.GetString("app.time-format"), r.FormValue("createdAt"))

	if err != nil {
		createdAt = time.Now()
	}
	if last != nil && !last.CreatedAt.Before(createdAt) {
		return
	}

	u := &m.Unit{Title: r.FormValue("title"), Description: r.FormValue("description"), CreatedAt: createdAt}
	err = u.Insert(context.Background(), cm.DB, boil.Infer())
	cm.PanicIfNotNil(err)

	res, errM := json.Marshal(cm.PostResult{Id: u.UnitID})
	cm.PanicIfNotNil(errM)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	removeEarlyIfNeed()
}

func removeEarlyIfNeed() {
	count, err := m.Units().Count(context.Background(), cm.DB)
	cm.PanicIfNotNil(err)

	if count < viper.GetInt64("app.MAX_LIMIT_UNITS") {
		return
	}
	border, err := m.Units(OrderBy(`created_at desc`), Offset(viper.GetInt("app.BASE_LIMIT_UNITS"))).One(context.Background(), cm.DB)
	cm.PanicIfNotNil(err)

	units, err := m.Units(Where("created_at < ?", border.CreatedAt)).All(context.Background(), cm.DB)
	cm.PanicIfNotNil(err)

	_, err = units.DeleteAll(context.Background(), cm.DB)
	cm.PanicIfNotNil(err)
}
