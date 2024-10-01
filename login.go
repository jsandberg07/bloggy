package main

// handle of all commands
// func handlerCommand(s *state, cmd command) error {}

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires username")
	}

	s.cfg.SetUser(cmd.args[0])
	fmt.Println("Login successful. User has been set.")
	return nil
}
