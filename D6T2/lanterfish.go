package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	days        int = 256
	maxDays     int = 8
	respawnDays int = 6
)

func main() {
	memory := make(map[int]int)
	input := strings.Split(os.Args[1], ",")
	for _, s := range input {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "lanterfisch: loading numbers %v\n", err)
			os.Exit(1)
		}
		memory[val]++
	}
	for i := 0; i < days; i++ {
		newFishes := memory[0]
		for j := 0; j < maxDays; j++ {
			memory[j] = memory[j+1]
		}
		memory[respawnDays] += newFishes
		memory[maxDays] = newFishes
	}
	counter := 0
		for _, v := range memory {
			if v > 0 {
				counter+=v
			}
		}
		fmt.Printf(" -> %v\n", counter)
}
