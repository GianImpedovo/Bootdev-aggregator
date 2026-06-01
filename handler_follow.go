package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GianImpedovo/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("Insert a url to follow")
	}

	argURL := cmd.arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), argURL)
	if err != nil {
		return errors.New("No feed")
	}
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return errors.New("No user")
	}

	feedFollowCreated, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return errors.New("Can't create feed follow")
	}

	fmt.Printf("%s\n", feedFollowCreated.UserName)
	fmt.Printf("%s\n", feedFollowCreated.FeedName)

	return nil
}
