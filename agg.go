package main

import (
	"bloggy/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
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

	return nil
}

// fetch rss feed, parse, print to console in a long running loop
func scrapeFeeds(s *state) error {
	// get the next feed to fetch from DB
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return errors.New("Error getting next feed to fetch.")
	}

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
	// iterate over the items in the feed and print their titles
	fmt.Println(rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
