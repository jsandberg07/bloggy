package main

import (
	"bloggy/internal/database"
	"internal/config"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
