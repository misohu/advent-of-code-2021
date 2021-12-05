package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dive: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	horizont := 0
	depth := 0
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		amount, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "dive: Error during reading %v", err)
			os.Exit(1)
		}
		switch words[0] {
		case "forward":
			horizont += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	f.Close()
	fmt.Printf("HORIZONT: %v, DEPTH %v, RESULT %v\n", horizont, depth, horizont*depth)
}
