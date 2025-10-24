package main

import (
	"context"
	"fmt"
)

func userHandler(s *state, cmd command) error {
	userList, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range userList {
		if user.Name == s.config.Current_user_name {
			fmt.Printf("%s (current)", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
