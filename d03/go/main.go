package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func countCol(in []string, pos int) (x0, x1 int) {
	for _, num := range in {
		if string(num[pos]) == "0" {
			x0 += 1
		} else {
			x1 += 1
		}
	}
	return
}

func mostCommonBit(in []string, pos int, tie string) string {
	x0, x1 := countCol(in, pos)
	if x0 > x1 {
		return "0"
	} else if x0 < x1 {
		return "1"
	} else {
		return tie
	}
}

func leastCommonBit(in []string, pos int, tie string) string {
	x0, x1 := countCol(in, pos)
	if x0 > x1 {
		return "1"
	} else if x0 < x1 {
		return "0"
	} else {
		return tie
	}
}

func filterBit(in []string, pos int, val string) (filtered []string) {
	for _, num := range in {
		if string(num[pos]) == val {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

func solveP1(in []string) int64 {
	sums := make([]int, len(in[0]))
	for _, num := range in {
		for i, c := range num {
			v, _ := strconv.Atoi(string(c))
			sums[i] += v
		}

	}
	gamma := ""
	epsilon := ""
	for _, sum := range sums {
		if sum > len(in)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}

func solveP2(in []string) int64 {
	pos := 0
	f := in
	for {
		most := mostCommonBit(f, pos, "1")
		f = filterBit(f, pos, most)
		if len(f) == 1 {
			break
		}
		pos++
	}
	o2, _ := strconv.ParseInt(f[0], 2, 64)

	pos = 0
	f = in
	for {
		least := leastCommonBit(f, pos, "0")
		f = filterBit(f, pos, least)
		if len(f) == 1 {
			break
		}
		pos++
	}
	co2, _ := strconv.ParseInt(f[0], 2, 64)

	return o2 * co2
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
	fmt.Printf("p1: %d\n", solveP1(in))
	fmt.Printf("p2: %d\n", solveP2(in))
}
