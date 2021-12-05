package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	windowSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "increased_wind: wrong window size %v", err)
		os.Exit(1)
	}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "increased_wind: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	prev := -1
	counter := 0
	buffer := []int{}
	for scanner.Scan() {
		new, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "increased: Error during reading %v", err)
			os.Exit(1)
		}
		buffer = append(buffer, new)
		if len(buffer) < windowSize {
			continue
		}
		newSum := sum(buffer)
		if prev != -1 {
			if newSum > prev {
				counter++
			}
		}
		prev = newSum
		buffer = buffer[1:]
	}
	f.Close()
	fmt.Println(counter)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
