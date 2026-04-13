package main

import (
	"errors"
	"context"
	"log"
	"time"
	"fmt"
	"github.com/JuanasoKsKs/agregator/internal/database"
	"github.com/google/uuid"
	"database/sql"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("The AddFeed handler expects a Two arguments, the Name and the URL")
	}
	ctx := context.Background()
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

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedFetched(ctx)
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return err
	}
	fmt.Println("Fetching New feed...........")
	Rssfeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}
	fmt.Printf("============   Saving posts from: (%s)  ============\n", Rssfeed.Channel.Title)
	for _, item := range Rssfeed.Channel.Item {
		// layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		// layout == time.RFC1123Z
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}
		null_description := sql.NullString{
			String: item.Description,
			Valid: item.Description != "",
		}
		params := database.CreatePostParams{
			ID          : uuid.New(),
			CreatedAt   : time.Now(),
			UpdatedAt   : time.Now(),
			Title       : item.Title,
			Url         : item.Link,
			Description : null_description,
			PublishedAt : t,
			FeedID      : nextFeed.ID,
		}
		post, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			return err
		}
		fmt.Println(post.Title)

	}
	return nil
}