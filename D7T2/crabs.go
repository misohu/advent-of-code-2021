package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := os.Args[1]
	starts := []int{}
	min := int(^uint(0) >> 1)
	max := -1
	for _, s := range strings.Split(inputs, ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Errorf("crabs: wrong input %v\n", err)
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		starts = append(starts, v)
	}
	dists := countDists(min, max)
	minDist := int(^uint(0) >> 1)
	for i := min; i <= max; i++ {
		s := 0
		failed := false
		for _, start := range starts {
			s += dists[int(math.Abs(float64(start - i)))]
			if s > minDist {
				failed = true
				break
			}
		}
		if failed {
			continue
		}
		if s < minDist {
			minDist = s
		}
	}
	fmt.Println(minDist)
}

func countDists(min, max int) map[int]int {
	dists := make(map[int]int)
	for i := min; i < max; i++ {
		temp := 1
		s := 0
		for j := i; j < max; j++ {
			s += temp
			temp++
		}
		dists[max-i] = s
	}
	return dists
}
