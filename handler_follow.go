package main

import (
	"errors"
	"context"
	"log"
	"time"
	"fmt"
	"github.com/JuanasoKsKs/agregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("The AddFeed Follow expects a One argument, the URL")
	}
	ctx := context.Background()
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		log.Fatalf("feed '%v' does NOT exists in the database: %v\n", url, err)
	}
	feed_follow_params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	row, err := s.db.CreateFeedFollow(ctx, feed_follow_params)
	if err != nil {
		log.Fatalf("Error creating feed_follow: %v\n", err)
	}
	fmt.Printf("the user: (%v) started following: (%v)\n", row.UserName, row.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("following command expects NO arguments")
	}
	followings, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatalf("Error getting the followings: %v\n", err)
	}
	fmt.Printf("%v is following: \n", user.Name)
	for _, feed := range followings {
		fmt.Println(feed.FeedName)
	}

	return nil
}