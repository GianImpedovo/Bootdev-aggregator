package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/GianImpedovo/aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return errors.New("Can't create feed follow")
	}

	fmt.Printf("%s:", user.Name)
	for _, u := range userFeeds {
		fmt.Printf("%s\n", u.FeedName)
	}

	return nil
}
