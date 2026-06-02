package main

import (
	"context"
	"errors"

	"github.com/GianImpedovo/aggregator/internal/database"
)

func handlerUnfollowing(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) == 0 {
		return errors.New("No url to eliminate the follow")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return errors.New("Feed not found")
	}

	err = s.db.DeleteFeedFollowByUserAndFeed(context.Background(), database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return errors.New("Can't delete de follow")
	}

	return nil
}
