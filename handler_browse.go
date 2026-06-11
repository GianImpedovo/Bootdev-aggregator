package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/GianImpedovo/aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2

	if len(cmd.arguments) == 1 {
		parseLimit, err := strconv.Atoi(cmd.arguments[0])

		if err != nil {
			return err
		}

		limit = parseLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Printf("- %s\n  URL: %s\n  Descripción: %s\n\n", p.Title, p.Url, p.Description)
	}

	return nil
}
