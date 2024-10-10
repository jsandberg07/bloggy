package main

import (
	"bloggy/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Agg takes a duration string for time between requests. ie 1s, 1m, 1h.")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		fmt.Println("Error turning duration string into interval.")
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	// loop
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
		log.Print("Scrape done.\n\n")
	}
}

// fetch rss feed, parse, print to console in a long running loop
func scrapeFeeds(s *state) error {
	// get the next feed to fetch from DB
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return errors.New("Error getting next feed to fetch.")
	}
	fmt.Printf("// Scraping %v\n", nextFeed.Url)

	// mark it as fetched
	mffParams := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}
	err = s.db.MarkFeedFetched(ctx, mffParams)
	// fetch the feed using the url
	rssFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return errors.New("Error fetching feed.")
	}

	// old code for just printing the posts, replaced with storing them
	// iterate over the items in the feed and print their titles
	/*
		fmt.Println(rssFeed.Channel.Title)
		for _, item := range rssFeed.Channel.Item {
			fmt.Println(item.Title)
		}
	*/

	err = storePosts(s, rssFeed, nextFeed.ID)
	if err != nil {
		return err
	}

	return nil
}

func storePosts(s *state, rssFeed *RSSFeed, feedID uuid.UUID) error {

	ctx := context.Background()
	timeLayouts := &[10]string{time.RFC1123, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123Z, time.RFC3339, time.ANSIC, time.UnixDate, time.RubyDate}

	for _, item := range rssFeed.Channel.Item {

		_, err := s.db.GetPostByURL(ctx, item.Link)
		if err == nil {
			// post was found, continue to next item
			continue
		}

		ct := pubTimeConverter(item.PubDate, timeLayouts)

		pubAt := sql.NullTime{
			Time:  ct,
			Valid: true,
		}

		cpParam := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubAt,
			FeedID:      feedID,
		}

		_, err = s.db.CreatePost(ctx, cpParam)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return err
			}
		}
	}

	return nil

}

func pubTimeConverter(t string, layouts *[10]string) time.Time {
	for i := range layouts {
		ct, err := time.Parse(layouts[i], t)
		if err == nil {
			return ct
		}
	}
	fmt.Println("Could not convert time. Using now as pubtime.")
	return time.Now()
}
