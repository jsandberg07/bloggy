package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {
	ctx := context.Background()
	currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	currentUserID := currentUser.ID

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

	feed, err := s.db.CreateFeed(ctx, newFeed)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
