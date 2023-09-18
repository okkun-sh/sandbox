package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file name")
		os.Exit(1)
		return
	}

	if err := run(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func run(path string) error {
	s, err := read(path)
	if err != nil {
		return err
	}

	fmt.Println(s)

	return nil
}

func read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", path, err)
	}

	return strings.TrimRight(string(data), "\n"), nil
}
