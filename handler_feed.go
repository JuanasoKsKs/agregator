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

	_, err = s.db.CreateFeed(ctx, feed_params)
	if err != nil {
		log.Fatalf("Error creating feed: \n", err)
	}
	fmt.Println("the feed has been added")
	return nil
}