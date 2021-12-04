package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Numbers  [][]string
	Marks    [][]bool
	Finished bool
}

func (c *Card) Mark(num string) bool {
	for i, row := range c.Numbers {
		for j, n := range row {
			if n == num {
				c.Marks[i][j] = true
				return true
			}
		}
	}
	return false
}

func (c *Card) Score() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !c.Marks[i][j] {
				v, _ := strconv.Atoi(c.Numbers[i][j])
				score += v
			}
		}
	}
	return score
}

func (c *Card) CheckRows() bool {
	for _, row := range c.Marks {
		for j, x := range row {
			if !x {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	return false
}

func (c *Card) CheckCols() bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !c.Marks[j][i] {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	return false
}

func (c *Card) Check() bool {
	return c.CheckRows() || c.CheckCols()
}

func NewCard() *Card {
	marks := make([][]bool, 0)
	for i := 0; i < 5; i++ {
		marks = append(marks, make([]bool, 5, 5))
	}
	return &Card{
		Numbers: make([][]string, 0),
		Marks:   marks,
	}
}

func playGame(numbers []string, cards []*Card) (string, *Card) {
	for _, num := range numbers {
		for _, card := range cards {
			if card.Mark(num) {
				if card.Check() {
					return num, card
				}

			}
		}
	}
	return "", &Card{} // should never happen
}

func loseGame(numbers []string, cards []*Card) (string, *Card) {
	wins := 0
	for _, num := range numbers {
		for _, card := range cards {
			if card.Finished {
				continue
			}
			if card.Mark(num) {
				if card.Check() {
					card.Finished = true
					wins++
					if wins == len(cards) {
						return num, card
					}
				}
			}
		}
	}
	return "", &Card{} // should never happen
}

func solveP1(in []string) int {
	numbers, cards := makeGame(in)
	lastNum, winCard := playGame(numbers, cards)
	winNum, _ := strconv.Atoi(lastNum)
	return winCard.Score() * winNum
}

func solveP2(in []string) int {
	numbers, cards := makeGame(in)
	lastNum, winCard := loseGame(numbers, cards)
	winNum, _ := strconv.Atoi(lastNum)
	return winCard.Score() * winNum
}

func makeGame(in []string) (numbers []string, cards []*Card) {
	numbers = strings.Split(in[0], ",")
	card := NewCard()
	for _, line := range in[2:] {
		if line == "" {
			cards = append(cards, card)
			card = NewCard()
			continue
		}
		card.Numbers = append(card.Numbers, strings.Fields(line))
	}
	cards = append(cards, card) // last one
	return numbers, cards
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
