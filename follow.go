package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Follow only takes a URL parameter.")
	}

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println("Couldn't get user while following a feed.")
		return err
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeed(ctx, url)

	cff := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(ctx, cff)
	if err != nil {
		fmt.Println("Problem creating Feed Follow")
		return err
	}

	fmt.Print(createdFeedFollow)
	return nil

}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("Following command doesn't take parameters.")
	}

	ctx := context.Background()
	fffu, err := s.db.GetFeedFollowsForUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		fmt.Print("Error getting feed follows for user")
		return err
	}

	for _, feed := range fffu {
		fmt.Println(feed.Name_2)
	}
	return nil
}
