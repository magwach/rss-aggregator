package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

func (apiCfg *apiConfig) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())

	feedsSlice := []Feed{}

	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("Failed to fetch all feeds: %v", err))
		return
	}

	for _, feed := range feeds {
		feedsSlice = append(feedsSlice, DatabaseFeedToFeed(feed))
	}

	RespondWithJson(w, 200, feedsSlice)
}

func (apiCfg *apiConfig) HandleFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing the body: %v", err))
		return
	}
	feed, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	RespondWithJson(w, 201, DatabaseFeedFollowingToFeedFollowing(feed))
}

func (apiCfg *apiConfig) HandleGetAllFollowingFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.GetAllFollowingFeeds(r.Context(), user.ID)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	followingFeeds := []FeedFollowing{}

	for _, feed := range feeds {
		followingFeeds = append(followingFeeds, DatabaseFeedFollowingToFeedFollowing(feed))
	}

	RespondWithJson(w, 201, followingFeeds)
}

func (apiCfg *apiConfig) HandleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowId := chi.URLParam(r, "feedId")

	feedId, err := uuid.Parse(feedFollowId)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing the feed Id: %v", err))
		return
	}
	err = apiCfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		UserID: user.ID,
		ID:     feedId,
	})

	if err == sql.ErrNoRows {
		fmt.Println("No matching record found.")
	} else if err != nil {
		fmt.Println("Database error:", err)
	} else {
		fmt.Println("Deleted:")
	}

	if err != nil {
		RespondWithError(w, 400, "Failed to unfollow feed")
		return
	}

	RespondWithJson(w, 200, "Unfollowed feed succesfully")
}

func (apiCfg *apiConfig) HandleGetPostsFromFollowing(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.GetAllPostsFromFollowingFeeds(r.Context(), user.ID)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error retrieving posts: %v", err))
		return
	}

	feedsSlice := []Post{}

	for _, feed := range feeds {
		feedsSlice = append(feedsSlice, DatabasePostToPost(feed))
	}

	RespondWithJson(w, 201, feedsSlice)
}
