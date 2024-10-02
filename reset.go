package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.ResetUser(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Users table reset.")
	return nil
}
