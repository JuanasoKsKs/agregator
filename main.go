package main
import(
	"github.com/JuanasoKsKs/agregator/internal/config"
	"log"
	"os"
	"database/sql"
	"github.com/JuanasoKsKs/agregator/internal/database"
	_ "github.com/lib/pq"
	//"fmt"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatalf("Specify a command\n")
	}
	cfgs, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	cmd := command{
		Name: arguments[1],
		Args: arguments[2:],
	}
	db, err := sql.Open("postgres", cfgs.DbURL)
	if err != nil {
		log.Fatalf("error establishing conection with the database: %v\n", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db: dbQueries,
		cfg: &cfgs,
	}
	cmds := commands{
		registeredCommands : make(map[string]func(*state, command) error),
	}
	
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerList)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("error running the command: %v\n", err)
	}
	

}