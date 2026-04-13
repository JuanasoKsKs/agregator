package main

import (
	"fmt"
	"errors"
	"context"
	"github.com/JuanasoKsKs/agregator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		errors.New("The Browse command expects None or One argument: The limit parameter")
	}
	var length int32 = 2
	if len(cmd.Args) == 1 {
		n, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return err
		}
		length = int32(n)
	}
	arguments := database.GetPostsForUserParams{
		UserID : user.ID,
		Limit  : length,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), arguments)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Title)
	}
	return nil
}