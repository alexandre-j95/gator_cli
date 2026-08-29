package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	commandmap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	input, ok := c.commandmap[cmd.name]
	if !ok {
		return errors.New("command doesn't exist")
	}
	return input(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandmap[name] = f
}
