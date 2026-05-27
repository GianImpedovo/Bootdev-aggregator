package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GianImpedovo/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {

	if len(cmd.arguments) < 2 {
		return errors.New("Forgot the name or the url of the feed")
	}

	fmt.Println(s.config.CurrentUserName)

	userId, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    userId.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
