package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/GianImpedovo/aggregator/internal/config"
	"github.com/GianImpedovo/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	var s state
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.config = &cfg
	db, err := sql.Open("postgres", s.config.DBURL)
	dbQueries := database.New(db)
	s.db = dbQueries

	clicmd := cliCommand{
		handlers: make(map[string]func(*state, command) error),
	}

	clicmd.register("login", handlerLogin)
	clicmd.register("register", handlerRegister)
	clicmd.register("reset", handlerReset)
	clicmd.register("users", handlerGetUsers)
	clicmd.register("agg", handlerAgg)
	clicmd.register("addfeed", handlerAddfeed)
	clicmd.register("feeds", handlerGetFeeds)

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println(errors.New("No enough arguments"))
		os.Exit(1)
	}

	cmd := command{
		name:      arguments[1],
		arguments: arguments[2:],
	}

	err = clicmd.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type cliCommand struct {
	handlers map[string]func(*state, command) error // "login": handlerLogin
}

func (c *cliCommand) run(s *state, cmd command) error {
	err := c.handlers[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *cliCommand) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
