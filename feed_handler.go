package main

import (
	"context"
	"fmt"
	"time"

	"github.com/akatakan/gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("feed name and url must be required")
	}
	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]
	
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("feed creation failed")
	}
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		FeedID:    feed.ID,
		UserID:    feed.UserID,
	})
	if err != nil {
		return fmt.Errorf("follow creation failed")
	}
	return nil
}

func listFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("cannot fetch feeds")
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("cannot fetch %s's feed", user.Name)
		}
		fmt.Printf("Feed Name:%s\n", feed.Name)
		fmt.Printf("Feed URL:%s\n", feed.Url)
		fmt.Printf("Created by:%s\n", user.Name)
	}
	return nil
}

func followFeed(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("feed url must be required")
	}
	feedUrl := cmd.arguments[0]
	feed, err := s.db.GetFeedFromUrl(ctx, feedUrl)
	if err != nil {
		return fmt.Errorf("feed not found")
	}
	createdFeed, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("feed follow canot created")
	}
	fmt.Printf("%s followed by %s\n", createdFeed.FeedName, createdFeed.UserName)
	return nil
}

func getFollowedFeeds(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	followedFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	for _, followedFeed := range followedFeeds {
		fmt.Println(followedFeed.FeedName)
	}
	if err != nil {
		return fmt.Errorf("user couldnt follow")
	}
	return nil
}

func unFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("feed url must be required")
	}
	feedUrl := cmd.arguments[0]
	ctx := context.Background()
	feed, err := s.db.GetFeedFromUrl(ctx, feedUrl)
	if err != nil {
		return err
	}
	s.db.DeleteFollowFromUser(ctx, database.DeleteFollowFromUserParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	fmt.Println("Unfollowed")
	return nil
}
