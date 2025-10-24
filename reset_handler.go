package main

import "context"

func resetHandler(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
