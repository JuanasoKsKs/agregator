package main

import (
	"fmt"
	"errors"
	"time"
	"context"
	"log"
	"github.com/JuanasoKsKs/agregator/internal/database"
	"github.com/google/uuid"
)

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




	err = s.cfg.SetUser(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the User has been set")


	return nil
}