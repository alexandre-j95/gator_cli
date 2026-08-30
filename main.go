package main

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/alexandre-j95/gator_cli/internal/config"
	"github.com/alexandre-j95/gator_cli/internal/database"
	"database/sql"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Couldn't read config file: %v", err)
		os.Exit(1)
	}
	
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Error opening database URL: %v\n", err)
		os.Exit(1)
	}

	s := state{database.New(db), &cfg}
	
	commands := commands{map[string]func(*state, command) error{}}
	
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)

	input := os.Args
	if len(input) < 2 {
		fmt.Println("Missing command")
		os.Exit(1)
	}

	commandStruct := command{input[1], input[2:]}
	err = commands.run(&s, commandStruct)
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
	
	os.Exit(0)
}
