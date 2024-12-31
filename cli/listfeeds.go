package cli

import (
	"context"
	"fmt"
	"log"

	
)

func HandleListFeeds(state *State, cmd Command) error {
	ctx := context.Background()
	
	feeds, err := state.DB.GetAllFeeds(ctx)
	if err != nil {
		log.Printf("Error fetching feeds: %v", err)
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Printf("- Name: %s\n  URL: %s\n  Created by: %s\n\n", 
			feed.FeedName, 
			feed.FeedUrl, 
			feed.UserName)
	}
	return nil
}
