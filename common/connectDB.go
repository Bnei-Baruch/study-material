package common

import (
	"database/sql"
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)

func Init() {
	var err error
	DB, err = sql.Open("postgres", viper.GetString("app.connection-string"))
	PanicIfNotNil(err)
}

func Close() {
	err := DB.Close()
	PanicIfNotNil(err)
}
