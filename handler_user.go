package main

import (
	"fmt"
	"errors"
	"context"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("The login hanler expects a single argument, the USERNAME")
	}

	username := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		log.Fatalf("user '%v' doesn't exists in the database: %v\n", username, err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the User has been set")
	return nil
}