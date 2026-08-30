package main

import (
	"github.com/alexandre-j95/gator_cli/internal/config"
	"github.com/alexandre-j95/gator_cli/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
