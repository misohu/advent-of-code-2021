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
		fmt.Fprintf(os.Stderr, "dive: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	counters := make(map[int]map[rune]int)
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			if counters[pos] == nil {
				counters[pos] = make(map[rune]int)
			}
			counters[pos][char]++
		}
	}
	gamma := ""
	epsilon := ""
	for i := 0; i < len(counters); i++ {
		if counters[i]['1'] > counters[i]['0'] {
			gamma += "1"
			epsilon += "0"
		} else if counters[i]['1'] < counters[i]['0'] {
			gamma += "0"
			epsilon += "1"
		} else {
			fmt.Fprint(os.Stderr, "binary_diagnostic: Unexpected problem")
			os.Exit(1)
		}
	}

	gammaVal, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary_diagnostic: %v", err)
	}

	epsilonVal, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary_diagnostic: %v", err)
	}
	fmt.Printf("Gamma: %s, Epsilon: %s\n", gamma, epsilon)
	fmt.Printf("Gamma: %v, Epsilon: %v\n", gammaVal, epsilonVal)
	fmt.Println(gammaVal * epsilonVal)
	f.Close()
}
