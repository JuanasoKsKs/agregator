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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return errors.New("The AddFeed handler expects a Two arguments, the Name and the URL")
	}
	current_user := s.cfg.CurrentUserName
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, current_user)
	if err != nil {
		log.Fatalf("user '%v' does NOT exists in the database: %v\n", current_user, err)
	}
	title := cmd.Args[0]
	url := cmd.Args[1]
	feed_params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: title,
		Url: url,
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feed_params)
	if err != nil {
		log.Fatalf("Error creating feed: %v\n", err)
	}
	fmt.Println("the feed has been added")

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

func handlerFeed(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
		return err
	}
	for i, feed := range feeds {
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("====== Feed %v ======\n", i + 1)
		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("URL: %v \n", feed.Url)
		fmt.Printf("User (author): %v\n", user.Name)

	}
	return nil
}