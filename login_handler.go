package main

import (
	"context"
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
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}
	config.SetUser(*s.config, user.Name)
	fmt.Printf("%s user has logged in\n", user.Name)
	return nil
}
