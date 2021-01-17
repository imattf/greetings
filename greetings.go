package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name given, return an error with a message.
	if name == "" {
		return "", errors.New("no name recieved")
	}

	// If a name is recieved, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome here!", name)
	return message, nil
}
