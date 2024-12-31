package cli

import (
	"context"
	"fmt"

	"github.com/polyfant/gator/internal/database"
)

func HandleUnfollow(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}

	url := cmd.Args[0]

	err := state.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Println("Successfully unfollowed feed")
	return nil
}