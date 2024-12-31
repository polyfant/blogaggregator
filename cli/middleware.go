package cli

import (
	

	"github.com/polyfant/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := GetAuthenticatedUser(s.DB)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}