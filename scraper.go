package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/samezio/rrs_aggregator/internal/database"
)

func scrapFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error in marking as fetched in DB", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error in fetching rss feed", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldn't parse date %v with err: %v\n", item.PubDate, err)
			return
		}
		if _, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		}); err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Println("Error in creating post", err)
			}
		}
	}
	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rssFeed.Channel.Item))
}

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration\n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		if feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency)); err != nil {
			log.Println("error in fetching feeds form DB", err)
		} else {
			wg := &sync.WaitGroup{}
			for _, feed := range feeds {
				wg.Add(1)
				go scrapFeed(wg, db, feed)
			}
			wg.Wait()
		}
	}
}
