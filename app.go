package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Usage: app <name>")
		os.Exit(1)
	}

	name := os.Args[1]

	if out, err := Dig(name); err != nil {
		fmt.Println(err)
		os.Exit(2)
	} else {
		fmt.Println(out)
	}
}

func Dig(arg string) (string, error) {
	fmt.Printf("Executing $ dig %s\n", arg)
	args := strings.Fields(arg)
	out, err := exec.Command("dig", args...).CombinedOutput()
	return string(out), err
}
