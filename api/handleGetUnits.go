package api

import (
	"context"
	"encoding/json"
	"github.com/Bnei-Baruch/study-material/common"
	"github.com/Bnei-Baruch/study-material/models"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
)

func handleGetUnits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	units, err := models.Units(qm.OrderBy(`created_at desc`), qm.Limit(viper.GetInt("app.get-limit"))).All(context.Background(), common.DB)
	common.PanicIfNotNil(err)

	var forJson []common.UnitForClient
	for _, u := range units {
		forJson = append(forJson, common.UnitForClient{Title: u.Title, Description: u.Description})
	}

	b, err := json.Marshal(forJson)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
