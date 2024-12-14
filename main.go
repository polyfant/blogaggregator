package main

import (
	"fmt"
	"os"

	"github.com/polyfant/gator/internal/config"
	"github.com/polyfant/gator/cli"
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

	// Initialize application state
	state := &cli.State{
		Config: cfg,
	}

	// Initialize commands
	commands := cli.NewCommands()
	
	// Register command handlers
	commands.Register("login", cli.HandleLogin)

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