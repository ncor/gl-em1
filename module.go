package example_module

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hi %v. Welcome!", name)
}
