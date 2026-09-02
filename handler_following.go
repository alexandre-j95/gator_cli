package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	u, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), u.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}
	fmt.Printf("The user %s is currently following:\n", u.Name)
	for _, item := range follows {
		fmt.Printf(" > %s\n", item.FeedName)
	}
	fmt.Println("===============")
	return nil
}
