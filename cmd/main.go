package main

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-cobra-cli-sample/internal/command"
)

func main() {
	err := command.Execute()
	if err != nil {
		log.Fatalf("command.Execute error: %v", err)
	}
}
