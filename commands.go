package main

import (
	"log"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		log.Fatalf("the command is not in the commands map")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.registeredCommands[name] = f
	return nil
}