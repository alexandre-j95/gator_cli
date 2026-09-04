package main

import (
	"context"
	"errors"

	"github.com/alexandre-j95/gator_cli/internal/database"
)

func handlerUnfollow(s* state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("must specify feed to unfollow")
	}

	feed, err :=s.db.GetFeedFromURL(context.Background(), cmd.args[0])
	if err != nil { return err }
	err = s.db.DeleteFollow(context.Background(), database.DeleteFollowParams{user.ID, feed.ID})
	if err != nil { return err }
	return nil
}
