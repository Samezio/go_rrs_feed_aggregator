package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/samezio/rrs_aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error in parsing JSON: %v", err))
		return
	} else if user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	} else {
		responseWithJSON(w, 200, databaseUserToUser(user))
	}
}
func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	if posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	}); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't fetch posts: %v\n", err))
	} else {
		responseWithJSON(w, 200, databasePostsToPosts(posts))
	}
}
