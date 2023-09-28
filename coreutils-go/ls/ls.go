package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func run() error {
	filenames, err := list()
	if err != nil {
		return err
	}

	maxLen := 0
	for _, name := range filenames {
		if maxLen < len(name) {
			maxLen = len(name)
		}
	}

	for _, name := range filenames {
		padding := maxLen + 1 - len(name)
		fmt.Printf("%s%*s", name, padding, " ")
	}
	fmt.Println()

	return nil
}

func list() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, len(files))
	for i, file := range files {
		filenames[i] = file.Name()
	}

	return filenames, nil
}
