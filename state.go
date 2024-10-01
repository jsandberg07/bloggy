package main

import "internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
