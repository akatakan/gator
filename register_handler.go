package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/akatakan/gator/internal/config"
	"github.com/akatakan/gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}
	if len(cmd.arguments) != 1 {
		return errors.New("just give username")
	}
	username := cmd.arguments[0]
	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	err = config.SetUser(*s.config, username)
	if err != nil {
		return err
	}
	fmt.Printf("%s was created\n", username)
	return nil
}
