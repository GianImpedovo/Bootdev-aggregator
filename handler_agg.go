package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/GianImpedovo/aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) != 1 {
		return errors.New("No arguments")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func scrapeFeeds(s *state) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, v := range rssFeed.Channel.Item {
		actualTime, err := time.Parse(time.RFC1123Z, v.PubDate)
		if err != nil {
			return err
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       v.Title,
			Url:         v.Link,
			Description: v.Description,
			PublishedAt: actualTime,
			FeedID:      feed.ID,
		})

		if err != nil {
			// Si el error es simplemente porque el post ya existe, lo ignoramos silenciosamente
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			// Si es cualquier otro error, lo reportamos en la consola para saber qué falló
			fmt.Printf("Error al crear el post: %v\n", err)
			continue
		}

		// fmt.Printf("%s\n", v.Title)
	}

	return nil
}
