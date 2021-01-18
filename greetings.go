package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name given, return an error with a message.
	if name == "" {
		return "", errors.New("no name recieved")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

	// If a name is recieved, return a value that embeds the name
	// in a greeting message.
	// message := fmt.Sprintf("Hi, %v. Welcome here!", name)
	// return message, nil
}

// init sets initial values for vairable used in the function.
func init() {
	// For truly random greetings, import "time" and replace the call
	// to rand.Seed with:
	//
	// rand.Seed(time.Now().UnixNano())
	//
	// Callling rand.Seed with a constant value means that we always
	// generate the same psuedo-random sequence.
	rand.Seed(1)
}

// randomFormat returns one set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"All Hail, %v! Well met!",
	}

	// Return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}
