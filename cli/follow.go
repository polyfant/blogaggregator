package cli

import (
	"context"
	"fmt"
	"github.com/polyfant/gator/internal/database"
)

func HandleFollow(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}

	url := cmd.Args[0]
	
	// Get the feed by URL
	feed, err := state.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error finding feed: %w", err)
	}

	// Create the feed follow
	feedFollow, err := state.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("You (%s) are now following '%s'\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func HandleFollowing(state *State, cmd Command, user database.User) error {
	// Get all feed follows for user
	follows, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You are not following any feeds")
		return nil
	}

	fmt.Println("You are following:")
	for _, follow := range follows {
		fmt.Printf("- %s\n", follow.FeedName)
	}
	return nil
}