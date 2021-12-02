package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	Dir  string
	Dist int
}

func solveP1(in []Movement) int {
	x, y := 0, 0
	for _, move := range in {
		switch move.Dir {
		case "forward":
			x += move.Dist
		case "down":
			y += move.Dist
		case "up":
			y -= move.Dist
		}
	}
	return x * y
}

func solveP2(in []Movement) int {
	x, y, aim := 0, 0, 0
	for _, move := range in {
		switch move.Dir {
		case "forward":
			x += move.Dist
			y += aim * move.Dist
		case "down":
			aim += move.Dist
		case "up":
			aim -= move.Dist
		}
	}
	return x * y
}

func readInput(in io.Reader) (data []Movement) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		move := strings.Split(s.Text(), " ")
		dist, _ := strconv.Atoi(move[1])
		data = append(data, Movement{move[0], dist})
	}
	return data
}

func main() {
	in := readInput(os.Stdin)
	fmt.Printf("p1: %d\n", solveP1(in))
	fmt.Printf("p2: %d\n", solveP2(in))
}
