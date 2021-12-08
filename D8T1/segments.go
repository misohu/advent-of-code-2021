package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var simples = []int{2, 4, 3, 7}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "segments: wrong input file %v", err)
		os.Exit(1)
	}
	inputs := readFile(f)
	counter := 0
	for _, input := range inputs {
		for _, s := range input["result"] {
			for _, l := range simples {
				if len(s) == l {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
}

func readFile(f *os.File) []map[string][]string {
	inputs := []map[string][]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " | ")
		input := map[string][]string{
			"start":  strings.Split(lineParts[0], " "),
			"result": strings.Split(lineParts[1], " "),
		}
		inputs = append(inputs, input)
	}
	return inputs
}
