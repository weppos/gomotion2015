package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Usage: app <name>")
		os.Exit(1)
	}

	name := os.Args[1]
	var out string
	var err error

	out, err = Dig(name)
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
