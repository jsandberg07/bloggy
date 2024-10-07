package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command, user database.User) error {

	currentUserID := user.ID

	if len(cmd.args) < 2 {
		return errors.New("Feed creation requires name and URL")
	}

	feedName := cmd.args[0]
	url := cmd.args[1]

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    currentUserID,
	}

	ctx := context.Background()
	feed, err := s.db.CreateFeed(ctx, newFeed)
	if err != nil {
		return err
	}

	// create a feed follow record for the current user when they add a feed
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(ctx, newFeedFollow)
	if err != nil {
		fmt.Println("Error creating feed follow when adding feed for user.")
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil

}
