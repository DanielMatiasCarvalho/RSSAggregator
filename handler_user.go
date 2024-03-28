package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DanielMatiasCarvalho/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	error := decoder.Decode(&params)
	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	user, error := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", error))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
