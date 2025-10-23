package main

import (
	"errors"
	"fmt"

	"github.com/akatakan/gator/internal/config"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}
	if len(cmd.arguments) != 1 {
		return errors.New("just give username")
	}
	username := cmd.arguments[0]
	err := config.SetUser(*s.config, username)
	if err != nil {
		return err
	}
	fmt.Printf("%s user has been set\n", username)
	return nil
}
