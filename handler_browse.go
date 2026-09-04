package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alexandre-j95/gator_cli/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		if len(cmd.args) == 1 {
			specifiedLimit, err := strconv.Atoi(cmd.args[0])
			if err != nil {
				return fmt.Errorf("invalid limit: %w", err)
			}
			limit = specifiedLimit
		}
	}

	postList, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{user.ID, int32(limit)})
	if err != nil {
		return fmt.Errorf("unable to retrieve user's posts: %w", err)
	}
	for _, item := range postList {
		fmt.Printf("--- %s ---\n", item.Title)
		fmt.Printf("Link:        %s\n", item.Url)
		if item.Description.Valid && item.Description.String != "" {
			fmt.Printf("Description: %s\n", item.Description.String)
		}
		if item.PublishedAt.Valid {
			fmt.Printf("Published:   %v\n", item.PublishedAt.Time.Format("Mon Jan 2"))
		}
		fmt.Println()
	}

	return nil
}
