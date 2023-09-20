package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := flag.Bool("n", false, "number the output lines")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Missing file name")
		os.Exit(1)
		return
	}

	if err := run(flag.Args()[0], *n); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func run(path string, b bool) error {
	s, err := read(path, b)
	if err != nil {
		return err
	}

	fmt.Print(s)

	return nil
}

func read(path string, n bool) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", path, err)
	}

	var content string
	if n {
		var builder strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		num := 1
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintf(&builder, "%6d\t%s\n", num, line)
			num++
		}
		content = builder.String()
	} else {
		content = string(data)
	}

	return content, nil
}
