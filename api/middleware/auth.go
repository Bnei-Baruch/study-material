package middleware

import (
	"context"
	"github.com/coreos/go-oidc"
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

			if rsa := fromAppScript(r); rsa != "" {
				log.Printf("RSA Authorization %s", rsa)
				if ok := validateRsa(rsa); ok {
					next.ServeHTTP(w, r)
					return
				}

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

func validateRsa(rsa string) bool {
	return true
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
	rsaHeader := strings.Split(strings.TrimSpace(r.Header.Get("Authorization")), " ")

	log.Printf("Authorization header from app script %s", rsaHeader)
	if len(rsaHeader) == 2 &&
		strings.ToLower(rsaHeader[0]) == "rsa" &&
		len(rsaHeader[1]) > 0 {
		return rsaHeader[1]
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
