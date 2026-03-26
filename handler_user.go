package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("The login hanler expects a single argument, the USERNAME")
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the User has been set")
	return nil
}