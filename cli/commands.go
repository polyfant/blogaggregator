package cli

import (
	"fmt"
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
