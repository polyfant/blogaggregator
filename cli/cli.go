package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/polyfant/gator/internal/config"
	"github.com/polyfant/gator/internal/database"
)

// State holds the application state
type State struct {
	Config *config.Config
	DB     *database.Queries
}

// Command represents a CLI command with its arguments
type Command struct {
	Name string
	Args []string
}

func HandleLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("usage: login <username>")
	}
	username := cmd.Args[0]

	// Check if user exists in database
	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %s does not exist", username)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to update config: %v", err)
	}

	fmt.Printf("Logged in as %s\n", username)
	fmt.Printf("User details: %+v\n", user)
	return nil
}

func HandleRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("usage: register <username>")
	}
	username := cmd.Args[0]

	// Check if user already exists
	_, err := s.DB.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("user %s already exists", username)
	}

	// Create new user
	now := time.Now().UTC()
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Update config
	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to update config: %v", err)
	}

	fmt.Printf("Created user %s\n", username)
	fmt.Printf("User details: %+v\n", user)
	return nil
}

func HandleReset(s *State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("Failed to reset users:", err)
		return err
	}
	fmt.Println("Successfully deleted all users")
	return nil
}

func GetAuthenticatedUser(db *database.Queries) (database.User, error) {
	cfg, err := config.Read()
	if err != nil {
		return database.User{}, fmt.Errorf("error reading config: %v", err)
	}

	if cfg.CurrentUserName == "" {
		return database.User{}, fmt.Errorf("no user logged in")
	}

	return db.GetUser(context.Background(), cfg.CurrentUserName)
}
