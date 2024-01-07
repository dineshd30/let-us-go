package main

import (
	"fmt"
	"strings"
)

// Create function type Message
type Message func(name string) string

// Create toLower receiver function
func (m Message) toLower(name string) string {
	return strings.ToLower(m(name))
}

// Create toUpper receiver function
func (m Message) toUpper(name string) string {
	return strings.ToUpper(m(name))
}

func main() {
	// Create instance of Message function by passing function literal
	msg := Message(func(name string) string {
		return "Hey, " + name
	})

	fmt.Println(msg("Dinesh"))
	fmt.Println(msg.toUpper("Dinesh"))
	fmt.Println(msg.toLower("Dinesh"))
}
