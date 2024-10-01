package main

import "errors"

type commands struct {
	// map that takes a string and returns a function (high or first order or whatever)
	cmdsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmdsMap[cmd.name]
	if !ok {
		return errors.New("Command not found.")
	}

	err := f(s, cmd)
	return err

}
