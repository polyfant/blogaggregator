package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/polyfant/gator/cli"
	"github.com/polyfant/gator/internal/config"
	"github.com/polyfant/gator/internal/database"
)

func main() {
	// Check for minimum number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [arguments]")
		os.Exit(1)
	}

	// Load config
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	// Open database connection
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error pinging database: %v\n", err)
		os.Exit(1)
	}

	// Initialize application state
	state := &cli.State{
		DB:     database.New(db),
		Config: cfg,
	}

	// Initialize commands
	commands := cli.NewCommands()
	
	// Register command handlers
	commands.Register("login", cli.HandleLogin)
	commands.Register("register", cli.HandleRegister)
	commands.Register("reset", cli.HandleReset)
	commands.Register("users", cli.HandleUsers)
	commands.Register("agg", cli.HandleAgg)
	commands.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandleAddFeed))
	commands.Register("feeds", cli.HandleListFeeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(cli.HandleFollow))
	commands.Register("following", cli.MiddlewareLoggedIn(cli.HandleFollowing))
	commands.Register("unfollow", cli.MiddlewareLoggedIn(cli.HandleUnfollow))
	commands.Register("browse", cli.MiddlewareLoggedIn(cli.HandleBrowse))

	// Create command from arguments
	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// Run the command
	if err := commands.Run(state, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}