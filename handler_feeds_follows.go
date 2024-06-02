package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/samezio/rrs_aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error in parsing JSON: %v", err))
		return
	} else if feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed follow : %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseFeedFollowToFeedFollow(feedFollow))
	}
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	if feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows from DB: %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
	}
}

func (apiCfg *apiConfig) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	if feedFollowIDString := chi.URLParam(r, "feedFollowID"); feedFollowIDString == "" {
		responseWithError(w, 400, "Feed id not found in Url")
		return
	} else if feedFollowID, err := uuid.Parse(feedFollowIDString); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error in parsing UUID: %v", err))
		return
	} else if err := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows from DB: %v", err))
		return
	} else {
		responseWithJSON(w, 200, struct{}{})
	}
}
