package cli

import (
	"context"
	"fmt"

	"github.com/polyfant/gator/internal/database"
)

func HandleAddFeed(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := state.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	// Create a feed follow
	_, err = state.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
    	UserID: user.ID,
    	FeedID: feed.ID,
	})
	if err != nil {
    	return fmt.Errorf("error following feed: %w", err)
}

	fmt.Printf("Feed created successfully:\n")
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("URL: %v\n", feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)
	fmt.Printf("Created At: %v\n", feed.CreatedAt)

	return nil
}
