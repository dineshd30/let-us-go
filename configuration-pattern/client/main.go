package main

import (
	"log"

	"github.com/dineshd30/golang/configuration-pattern/server"
)

func main() {
	// Server with default configuration
	server1 := server.New()
	log.Printf("server1 - %v", server1)

	// Server with custom max connections
	server2 := server.New(server.WithMaxCon(100))
	log.Printf("server2 - %v", server2)

	// Server with custom tls
	server3 := server.New(server.WithTLS)
	log.Printf("server3 - %v", server3)

	// Server with custom id
	server4 := server.New(server.WithId("Custom"))
	log.Printf("server4 - %v", server4)

	// Server with custom max connections and tls
	server5 := server.New(server.WithMaxCon(50), server.WithTLS)
	log.Printf("server5 - %v", server5)

}
