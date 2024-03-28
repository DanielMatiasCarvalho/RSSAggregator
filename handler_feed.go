package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DanielMatiasCarvalho/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:name`
		Url  string `json:url`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	error := decoder.Decode(&params)
	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	feed, error := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", error))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, error := apiCfg.DB.GetFeeds(r.Context())
	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", error))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
