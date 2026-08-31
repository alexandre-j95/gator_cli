package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/alexandre-j95/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("missing arguments")
	}

	currentUserName := s.cfg.CurrentUserName
	currentUserStruct, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
			Url:       cmd.args[1],
			UserID:    currentUserStruct.ID,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Feed id: %v\n", feed.ID)
	log.Printf("Feed created at: %v\n", feed.CreatedAt)
	log.Printf("Feed updated at: %v\n", feed.UpdatedAt)
	log.Printf("Feed name: %v\n", feed.Name)
	log.Printf("Feed url: %v\n", feed.Url)
	log.Printf("Feed created by UID: %v\n", feed.UserID)
	return nil
}
