package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Dir    rune
	Amount int
}

type Point struct {
	X, Y int
}

func ReadFromFile(filePath string) ([]Command, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	res := []Command{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		conv, _ := strconv.Atoi(split[1])
		res = append(res, Command{Dir: rune(split[0][0]), Amount: conv})
	}
	return res, nil
}

func AdjustTail(head Point, tail *Point, visitedPos *map[Point]bool, isLast bool) {
	if head.X == tail.X {
		if head.Y > tail.Y && Abs(head.Y-tail.Y) >= 2 {
			(*tail).Y++
		} else if tail.Y-head.Y >= 2 {
			(*tail).Y--
		}
	} else if head.Y == tail.Y {
		if head.X > tail.X && Abs(head.X-tail.X) >= 2 {
			(*tail).X++
		} else if tail.X-head.X >= 2 {
			(*tail).X--
		}
	} else {
		shouldMove := Abs(tail.X-head.X) >= 2 || Abs(tail.Y-head.Y) >= 2
		if head.X < tail.X && shouldMove {
			(*tail).X--
		} else if shouldMove {
			(*tail).X++
		}

		if head.Y < tail.Y && shouldMove {
			(*tail).Y--
		} else if shouldMove {
			(*tail).Y++
		}
	}
	if isLast {
		(*visitedPos)[*tail] = true
	}
}

func AdjustTails(head Point, tails []*Point, visitedPos *map[Point]bool) {
	AdjustTail(head, tails[0], visitedPos, false)
	for idx := 1; idx < len(tails); idx++ {
		AdjustTail(*tails[idx-1], tails[idx], visitedPos, idx == len(tails)-1)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solve(comms []Command) int {
	visitedPos := map[Point]bool{}
	head := Point{X: 0, Y: 0}
	tails := []*Point{}
	for idx := 0; idx < 9; idx++ {
		tails = append(tails, &Point{X: 0, Y: 0})
	}
	visitedPos[*tails[0]] = true

	for _, cmd := range comms {
		for st := 0; st < cmd.Amount; st++ {
			switch cmd.Dir {
			case 'R':
				head.X++
				AdjustTails(head, tails, &visitedPos)
			case 'L':
				head.X--
				AdjustTails(head, tails, &visitedPos)
			case 'U':
				head.Y++
				AdjustTails(head, tails, &visitedPos)
			case 'D':
				head.Y--
				AdjustTails(head, tails, &visitedPos)
			}
		}
	}

	return len(visitedPos)
}

func main() {
	comms, err := ReadFromFile("./in-day9.txt")
	if err != nil {
		panic("Cannot read from file...")
	}

	fmt.Printf("%v\n", Solve(comms))
}
