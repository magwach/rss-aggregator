package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/magwach/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing the body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	RespondWithJson(w, 201, DatabaseFeedToFeed(feed))
}
