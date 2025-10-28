package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/magwach/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing the body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	RespondWithJson(w, 201, DatabaseUserToUser(user))
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJson(w, 200, DatabaseUserToUser(user))
}
