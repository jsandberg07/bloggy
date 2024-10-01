package main

// handle of all commands
// func handlerCommand(s *state, cmd command) error {}

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires username")
	}

	userName := cmd.args[0]

	// if given username doesnt exist, error
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	blankUser := database.User{}
	if user == blankUser {
		return errors.New("User not registered.")
	}

	s.cfg.SetUser(userName)
	fmt.Println("Login successful. User has been set.")
	return nil
}
