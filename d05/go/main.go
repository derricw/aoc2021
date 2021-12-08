package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

type Point [2]int

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p[0], p[1])
}

func NewPoint(s string) Point {
	v := strings.Split(s, ",")
	x, _ := strconv.Atoi(v[0])
	y, _ := strconv.Atoi(v[1])
	return Point{x, y}
}

type LineSegment struct {
	s Point
	d Point
}

func (l LineSegment) Horizontal() bool {
	return l.s[0] == l.d[0]
}

func (l LineSegment) Vertical() bool {
	return l.s[1] == l.d[1]
}

func (l LineSegment) Points() []Point {
	points := []Point{l.s, l.d}
	dx := l.d[0] - l.s[0]
	dy := l.d[1] - l.s[1]
	miny := min(l.d[1], l.s[1])
	minx := min(l.d[0], l.s[0])
	if dx == 0 {
		for i := miny + 1; i < miny+abs(dy); i++ {
			points = append(points, Point{l.s[0], i})
		}
	} else if dy == 0 {
		for i := minx + 1; i < minx+abs(dx); i++ {
			points = append(points, Point{i, l.s[1]})
		}
	} else { // 45 degrees
		slope := dy / dx
		if slope > 0 {
			for i := minx + 1; i < minx+abs(dx); i++ {
				points = append(points, Point{i, miny + (i - minx)})
			}
		} else {
			maxy := max(l.d[1], l.s[1])
			for i := minx + 1; i < minx+abs(dx); i++ {
				points = append(points, Point{i, maxy - (i - minx)})
			}
		}

	}
	return points
}

func NewLineSegment(line string) LineSegment {
	fields := strings.Split(line, " ")
	return LineSegment{
		s: NewPoint(fields[0]),
		d: NewPoint(fields[2]),
	}
}

func solveP1(in []string) int {
	segments := []LineSegment{}
	for _, l := range in {
		s := NewLineSegment(l)
		if s.Horizontal() || s.Vertical() {
			segments = append(segments, s)
		}
	}
	counter := map[Point]int{}
	for _, ls := range segments {
		for _, p := range ls.Points() {
			counter[p]++
		}
	}

	count := 0
	for _, c := range counter {
		if c > 1 {
			count++
		}
	}
	return count
}

func solveP2(in []string) int {
	segments := []LineSegment{}
	for _, l := range in {
		s := NewLineSegment(l)
		segments = append(segments, s)
	}
	counter := map[Point]int{}
	for _, ls := range segments {
		for _, p := range ls.Points() {
			counter[p]++
		}
	}

	count := 0
	for _, c := range counter {
		if c > 1 {
			count++
		}
	}
	return count
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
