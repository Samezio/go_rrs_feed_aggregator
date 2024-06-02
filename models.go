package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/samezio/rrs_aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, feed := range dbFeeds {
		feeds[i] = Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		}
	}
	return feeds
}
