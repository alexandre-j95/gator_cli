package main

import (
	"context"
	"fmt"

	"github.com/alexandre-j95/gator_cli/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}
	fmt.Printf("The user %s is currently following:\n", user.Name)
	for _, item := range follows {
		fmt.Printf(" > %s\n", item.FeedName)
	}
	fmt.Println("===============")
	return nil
}
