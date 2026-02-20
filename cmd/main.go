package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("AgentRegistry starting...")
	if err := initialize(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("AgentRegistry ready")
}

func initialize() error {
	return nil
}
