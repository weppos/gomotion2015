package main

import (
	"fmt"
	"log"
)

func main() {
	var out string
	var err error

	out, err = Dig("example.com")
	if err != nil {
		log.Fatalf("Execution failed: %s", err)
	} else {
		log.Println("Executed %s", out)
	}
}

func Dig(arg string) (string, error) {
	if (len(arg) < 5) {
		return "", fmt.Errorf("arg is too short")
	}

	return fmt.Sprintf("dig %s", arg), nil
}
