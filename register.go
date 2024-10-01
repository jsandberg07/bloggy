package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Can't create user without Username, or too many arguments added")
	}
	userName := cmd.args[0]
	// msg("Got username")

	cup := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}
	// msg("Created user params")

	user, err := s.db.CreateUser(context.Background(), cup)
	if err != nil {
		return err
	}
	// msg("Created user")

	// set the current user to the one that was registered
	s.cfg.SetUser(userName)

	fmt.Printf("User %v was created.\n", userName)
	fmt.Print(user)
	return nil
}
