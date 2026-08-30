package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alexandre-j95/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no argument given")
	}
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: cmd.args[0]})
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Error setting current user: %w", err)
	}

	log.Printf("User %s created successfuly", user.Name)
	return nil
}
