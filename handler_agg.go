package main

import (
	"context"
	"fmt"
	"errors"
	"time"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg command expects 1 argument: the TIME BETWEEN REPS")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("-----        Collecting feeds every: %s       ------\n", cmd.Args[0])
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C{
		scrapeFeeds(s)
	}


	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %+v\n", feed.Channel.Title)
	return nil
}