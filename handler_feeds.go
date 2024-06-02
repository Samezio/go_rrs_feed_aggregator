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
	} else if feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseFeedToFeed(feed))
	}
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	if feeds, err := apiCfg.DB.GetFeeds(r.Context()); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feeds from DB: %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseFeedsToFeeds(feeds))
	}
}
