package main

import (
	"fmt"
	"errors"
	"context"
	"log"
	"time"
	"github.com/JuanasoKsKs/agregator/internal/database"
	"github.com/google/uuid"
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

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("The register hanler expects a single argument, the USER")
	}
	username := cmd.Args[0]
	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	}
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if err == nil {
		log.Fatalf("user '%v' already exists in the database: %v\n", username, err)
	}

	user, err = s.db.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("Error creating user: \n", err)
	}
	//Set username =============
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the User has been set")
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		log.Fatalf("Error reseting the table: %v\n", err)
	}
	return nil
}

func handlerList(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		log.Fatalf("error getting the users: %v\n", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
		fmt.Println(user.Name)
		}
	}
	return nil
}

