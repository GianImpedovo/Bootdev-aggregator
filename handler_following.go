package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return errors.New("No user")
	}

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
