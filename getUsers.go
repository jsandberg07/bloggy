package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.cfg.CurrentUserName

	for u := range users {
		if users[u].Name == currentUser {
			fmt.Printf("* %v (current)\n", users[u].Name)
		} else {
			fmt.Printf("* %v\n", users[u].Name)
		}
	}
	return nil
}
