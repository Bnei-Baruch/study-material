package middleware

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(tokenVerifier *oidc.IDTokenVerifier) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/units" {
				next.ServeHTTP(w, r)
				return
			}

			if pass := fromAppScript(r); pass != "" {
				log.Printf("Authorization with passsword")
				if ok := validatePass(pass); ok {
					next.ServeHTTP(w, r)
					return
				}
				log.Printf("Authorization wrong passsword")
				authError(w, next)
				return
			}

			auth := parseToken(r)
			if auth == "" {
				authError(w, next)
			}

			token, err := tokenVerifier.Verify(context.TODO(), auth)
			if err != nil {
				authError(w, next)
			}
			if err := permission(token); err != nil {
				authError(w, next)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func validatePass(pass string) bool {
	return pass == viper.GetString("app.app-script-pass")
}

func permission(token *oidc.IDToken) error {
	var claims interface{}
	return token.Claims(&claims)
}

func authError(w http.ResponseWriter, h http.Handler) http.Handler {
	http.Error(w, "You are not authorized", http.StatusUnauthorized)
	return h
}

func fromAppScript(r *http.Request) string {
	passHeader := strings.Split(strings.TrimSpace(r.Header.Get("Authorization")), " ")

	if len(passHeader) == 2 && strings.ToLower(passHeader[0]) == "pass" && len(passHeader[1]) > 0 {
		log.Printf("Authorization header from app script")
		return passHeader[1]
	}
	return ""
}

func parseToken(r *http.Request) string {
	authHeader := strings.Split(strings.TrimSpace(r.Header.Get("Authorization")), " ")
	if len(authHeader) == 2 &&
		strings.ToLower(authHeader[0]) == "bearer" &&
		len(authHeader[1]) > 0 {
		return authHeader[1]
	}
	return ""
}
