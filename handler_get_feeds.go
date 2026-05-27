package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			continue
		}

		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}
