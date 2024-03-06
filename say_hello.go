package go_say_hello

import (
	"errors"
	"fmt"
	"math/rand"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(),name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hey bruh %v",
		"Haven't you had brekkie yet?",
		"JDON my soul %v '-')/",
	} 
	// Return a randomly selected message format by specifying
    // a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}