package main

import (
	"fmt"
	"internal/config"
	"os"
)

func main() {
	fmt.Println("Hello borld")

	configuration := config.Read()

	s := state{
		cfg: &configuration,
	}

	cmdMap := make(map[string]func(*state, command) error)
	cmds := commands{
		cmdsMap: cmdMap,
	}
	cmds.register("login", handlerLogin)

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

	err := cmds.run(&s, actualCommand)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
