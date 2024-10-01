package main

import (
	"bloggy/internal/database"
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
	dbQueries := database.New(db)
	s.db = dbQueries

	cmdMap := make(map[string]func(*state, command) error)
	cmds := commands{
		cmdsMap: cmdMap,
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
