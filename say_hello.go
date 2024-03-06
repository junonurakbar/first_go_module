package go_say_hello

import (
	"errors"
	"fmt"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Henlo Bruh %v '-')/", name)
	return message, nil
}