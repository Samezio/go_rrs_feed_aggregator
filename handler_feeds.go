package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/samezio/rrs_aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error in parsing JSON: %v", err))
		return
	} else if feed, err := apiCfg.DB.CreatFeed(r.Context(), database.CreatFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseFeedToFeed(feed))
	}
}

/*func (apiCfg *apiConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {
	if apiKey, err := auth.GetAPIKey(r.Header); err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	} else if user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Unable to find user: %v", err))
	} else {
		responseWithJSON(w, 200, databaseUserToUser(user))
	}
}*/
