package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	result, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error returning users from database: %w", err)
	}
	for _, person := range result {
		fmt.Printf("* %v", person.Name)
		if person.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}
