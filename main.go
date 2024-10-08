package main

import (
	"bloggy/internal/database"
	"context"
	"database/sql"
	"fmt"
	"internal/config"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello borld")

	configuration := config.Read()
	s := state{
		cfg: &configuration,
	}

	// fmt.Printf("// CHECK THIS URL: %v\n", s.cfg.DbURL)
	db, err := sql.Open("postgres", s.cfg.DbURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s.db = dbQueries

	cmdMap := make(map[string]func(*state, command) error)
	cmds := commands{
		cmdsMap: cmdMap,
	}

	// regular handler: func handlerCommand(s *state, cmd command) error {}
	// login required: func handlerCommand(s *state, cmd command, user database.User) error {}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Not enough arguments.")
		os.Exit(1)
	}

	// first argument is just the program name
	cmdName := arguments[1]
	cmdArgs := arguments[2:]
	actualCommand := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(&s, actualCommand)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("")
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
