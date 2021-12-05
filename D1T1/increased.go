package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "increased: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	counter := 0
	prev := -1
	for scanner.Scan() {
		new, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "increased: Error during reading %v", err)
			os.Exit(1)
		}
		if prev != -1 {
			if new > prev {
				counter++
			}
		}
		prev = new
	}
	f.Close()
	fmt.Println(counter)
}
