package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("空の名前")
	}
	return fmt.Sprintf("Hi %v. Welcome!", name), nil
}
