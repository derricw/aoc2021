package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func solveP1(in []int) int {
	lastDepth, increases := math.MaxInt, 0
	for _, depth := range in {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	return increases
}

func solveP2(in []int) int {
	lastSum, increases := math.MaxInt, 0
	for i := range in[:len(in)-2] {
		sum := in[i] + in[i+1] + in[i+2]
		if sum > lastSum {
			increases++
		}
		lastSum = sum
	}
	return increases
}

func readInput(in io.Reader) (data []int) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		d, _ := strconv.Atoi(s.Text())
		data = append(data, d)
	}
	return data
}

func main() {
	in := readInput(os.Stdin)
	fmt.Printf("p1: %d\n", solveP1(in))
	fmt.Printf("p2: %d\n", solveP2(in))
}
