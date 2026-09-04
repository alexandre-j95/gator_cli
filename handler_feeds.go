package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds to display")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s", feed.Name)
		fmt.Printf("URL: %s", feed.Url)
		fmt.Printf("Created by: %s", feed.Username.String)
	}
	return nil
}
