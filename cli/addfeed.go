package cli

import (
	"context"
	"fmt"
	"github.com/polyfant/gator/internal/database"
)

func HandleAddFeed(state *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get the current user
	user, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Create the feed
	feed, err := state.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	// Print the feed details
	fmt.Printf("Feed created successfully:\n")
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("Created at: %v\n", feed.CreatedAt)

	return nil
}
