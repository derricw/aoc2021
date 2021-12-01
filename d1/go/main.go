package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func solveP1(in []string) {
	lastDepth := math.MaxInt
	increases := 0
	for _, depthStr := range in {
		depth, _ := strconv.Atoi(depthStr)
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	fmt.Printf("p1: %d\n", increases)
}

func solveP2(in []string) {
	lastSum := math.MaxInt
	increases := 0
	for i := range in {
		if i+2 >= len(in) {
			break
		}
		d0, _ := strconv.Atoi(in[i])
		d1, _ := strconv.Atoi(in[i+1])
		d2, _ := strconv.Atoi(in[i+2])
		sum := d0 + d1 + d2
		if sum > lastSum {
			increases++
		}
		lastSum = sum
	}
	fmt.Printf("p2: %d\n", increases)

}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	return data
}

func main() {
	in := readInput(os.Stdin)

	solveP1(in)
	solveP2(in)
}
