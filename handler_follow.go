package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alexandre-j95/gator_cli/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("need feed url")
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to find feed: %w", err)
	}
	currentUserStruct, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FeedID:    feed.ID,
			UserID:    currentUserStruct.ID,
		})
	if err != nil {
		return err
	}

	fmt.Printf("User: %s is now following feed: %s", s.cfg.CurrentUserName, feed.Name)
	return nil
}
