package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/magwach/rss-aggregator/internal/database"
)

func scrape(
	db *database.Queries,
	concurrency int,
	interval time.Duration,
) {

	log.Printf("Scraping on %v goroutines every %v", concurrency, interval)

	ticker := time.NewTicker(time.Duration(interval))

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Printf("Error fectching feed %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeUrl(wg, db, feed)
		}
		wg.Wait()
	}

}

func scrapeUrl(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed as scraped: %v", feed.Name)
	}

	rssFeed, err := UrlToFeed(feed.Url)

	if err != nil {
		log.Printf("Error scraping %v, error: %v", feed.Name, err)
	}

	for _, item := range rssFeed.Channel.Item {
		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{Valid: item.Title == "", String: item.Title},
			Description: sql.NullString{Valid: item.Description == "", String: item.Description},
			PublishedAt: sql.NullTime{Valid: true, Time: item.PubDate.Time},
			Url:         feed.Url,
			FeedID:      feed.ID,
		})

		if err != nil {
			log.Printf("Error inserting %v to the database: %v", item.Title, err)
		}
	}

}
