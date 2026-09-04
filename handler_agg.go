package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Missing time interval")
	}
	
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	
	fmt.Println()
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)
	fmt.Println()
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil { return err }

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil { return err }

	rssfeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil { return err}

	for _, item := range rssfeed.Channel.Item {
		fmt.Printf("%+v\n", item.Title)
	}
	return nil
}
