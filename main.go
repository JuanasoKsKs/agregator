package main

import(
	"github.com/JuanasoKsKs/agregator/internal/config"
	"log"
	"os"
	//"fmt"
)

type state struct {
	cfg *config.Config
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatalf("Specify a command")
	}
	cfgs, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cmd := command{
		Name: arguments[1],
		Args: arguments[2:],
	}
	programState := &state{
		cfg: &cfgs,
	}
	cmds := commands{
		registeredCommands : make(map[string]func(*state, command) error),
	}
	
	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("error running the command: %v", err)
	}
	

}