package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/magwach/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}

type FeedFollowing struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowingToFeedFollowing(feed database.FeedFollowing) FeedFollowing {
	return FeedFollowing{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserId:    feed.UserID,
		FeedId:    feed.FeedID,
	}
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	Url         string     `json:"url"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func DatabasePostToPost(feed database.Post) Post {
	var description *string
	if feed.Description.Valid {
		description = &feed.Description.String
	}

	var published_at *time.Time
	if feed.PublishedAt.Valid {
		published_at = &feed.PublishedAt.Time
	}

	var title *string
	if feed.Title.Valid {
		title = &feed.Title.String
	}
	return Post{
		ID:          feed.ID,
		CreatedAt:   feed.CreatedAt,
		UpdatedAt:   feed.UpdatedAt,
		Title:       title,
		Description: description,
		PublishedAt: published_at,
		Url:         feed.Url,
		FeedID:      feed.FeedID,
	}
}
