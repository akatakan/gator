package main

import "errors"

func (c *commands) run(s *state, cmd command) error {
	executeCmd, ok := c.cmds[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return executeCmd(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
