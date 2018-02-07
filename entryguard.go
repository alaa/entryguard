package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	pattern := os.Getenv("PATTERN")
	key := os.Getenv("KEY")

	r, _ := regexp.Compile(pattern)
	if r.MatchString(key) {
		fmt.Println(os.Args[1:])
		command := os.Args[1:2]
		command_args := os.Args[2:]

		fmt.Println(command)
		fmt.Println(command_args)

		exec.Command(string(command[0]), command_args...).Output()
	} else {
		log.Fatal("Environment variables didn't match the define types")
	}
}
