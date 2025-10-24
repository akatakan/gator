package main

import (
	"context"
	"fmt"

	"github.com/akatakan/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.config.Current_user_name == "" {
			return fmt.Errorf("not logged in")
		}
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.config.Current_user_name)
		if err != nil {
			return fmt.Errorf("user not found")
		}
		return handler(s, cmd, user)
	}
}
