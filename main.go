package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/akatakan/gator/internal/config"
	"github.com/akatakan/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("cannot read config file")
	}
	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Println("cannot connect to db")
	}
	dbQueries := database.New(db)
	s := state{
		config: &cfg,
		db:     dbQueries,
	}
	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{
		cmdName,
		cmdArgs,
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
