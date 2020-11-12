package api

import (
	"context"
	"github.com/Bnei-Baruch/study-material/middleware"
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type App struct {
	verifier *oidc.IDTokenVerifier
	router   *mux.Router
	cors     *cors.Cors
}

func (a *App) Init() {
	a.initOidc()
	a.initCors()
	a.initRouters()

	handler := a.cors.Handler(
		middleware.AuthMiddleware(a.verifier)(
			a.router))

	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		log.Fatal(err)
	}
}

func (a *App) initOidc() {
	provider, err := oidc.NewProvider(context.TODO(), viper.GetString("app.provider-issuer"))
	if err != nil {
		log.Fatal("oidc.NewProvider", err)
	}

	a.verifier = provider.Verifier(&oidc.Config{
		SkipClientIDCheck: true,
	})
}
func (a *App) initCors() {
	a.cors = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type"},
		MaxAge:         0,
	})
}

func (a *App) initRouters() {
	a.router = mux.NewRouter()

	a.router.HandleFunc("/api/units", handleGetUnits).Methods(http.MethodGet)
	a.router.HandleFunc("/api/unit", handleAddUnit).Methods(http.MethodPost)
}
