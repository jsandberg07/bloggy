package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("Follow only takes a URL parameter.")
	}

	ctx := context.Background()

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

func handlerFollowing(s *state, cmd command, user database.User) error {
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	url := cmd.args[0]
	ctx := context.Background()
	uffuParams := database.UnfollowFeedForUserParams{
		Url:    url,
		UserID: user.ID,
	}

	err := s.db.UnfollowFeedForUser(ctx, uffuParams)
	if err != nil {
		return errors.New("Couldn't delete feed follow for user.")
	}

	fmt.Println("Feed unfollowed for user.")
	return nil
}
