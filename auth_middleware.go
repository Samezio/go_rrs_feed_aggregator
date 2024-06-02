package main

import (
	"fmt"
	"net/http"

	"github.com/samezio/rrs_aggregator/internal/auth"
	"github.com/samezio/rrs_aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiKey, err := auth.GetAPIKey(r.Header); err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		} else if user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey); err != nil {
			responseWithError(w, 400, fmt.Sprintf("Unable to find user: %v", err))
		} else {
			handler(w, r, user)
		}
	}
}
