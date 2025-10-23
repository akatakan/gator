package main

import (
	"github.com/akatakan/gator/internal/config"
	"github.com/akatakan/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}
