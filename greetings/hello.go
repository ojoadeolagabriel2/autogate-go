package greetings

import "fmt"

func SayHello(message string) string {
	return fmt.Sprintf("Hi, %v. welcome", message)
}
