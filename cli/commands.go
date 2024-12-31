package cli

import (
	"context"
	"fmt"
	"time"
	"log"

	"github.com/polyfant/gator/internal/database"
	"github.com/polyfant/gator/internal/rss"
)

type CommandHandler func(*State, Command) error

type Commands struct {
	handlers map[string]CommandHandler
}

func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]CommandHandler),
	}
}

func (c *Commands) Register(name string, handler CommandHandler) {
	c.handlers[name] = handler
}

func (c *Commands) Run(state *State, cmd Command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(state, cmd)
}

func HandleUsers(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

func HandleAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s, timeBetweenRequests.String())
	}
}

func scrapeFeeds(s *State, interval string) {
	// Parse the duration string into minutes
	duration, err := time.ParseDuration(interval)
	if err != nil {
		log.Printf("Invalid interval format: %v", err)
		return
	}
	minutes := int64(duration.Minutes())
	
	feed, err := s.DB.GetNextFeedToFetch(context.Background(), minutes)
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.DB, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Items {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Items))
}