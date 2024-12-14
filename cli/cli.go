package cli

import (
	"fmt"
	"errors"

	"github.com/polyfant/gator/internal/config"
)


type State struct {
	Config *config.Config
}


type Command struct {
	Name string
	Args []string
}


type HandlerFunc func(s *State, cmd Command) error


type Commands struct {
	handlers map[string]func(*State, Command) error
}


func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]func(*State, Command) error),
	}
}


func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.handlers[name] = f
}


func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}


func HandleLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login command requires exactly one argument: username")
	}

	username := cmd.Args[0]
	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User set to: %s\n", username)
	return nil
}
