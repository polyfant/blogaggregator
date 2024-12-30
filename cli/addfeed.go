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

	user, err := GetAuthenticatedUser(state.DB)
	if err != nil {
		return err
	}

	feed, err := state.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed created successfully:\n")
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("URL: %v\n", feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)
	fmt.Printf("Created At: %v\n", feed.CreatedAt)

	return nil
}
