package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command interface {
	GetName() string
}

type Add struct {
	Args []int
}

func (a Add) GetName() string {
	return "addx"
}

type Noop struct{}

func (n Noop) GetName() string {
	return "noop"
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
		if split[0] == (Noop{}).GetName() {
			res = append(res, Noop{})
		} else {
			val, _ := strconv.Atoi(split[1])
			res = append(res, Add{Args: []int{val}})
		}
	}
	return res, nil
}

func Solve(cmds []Command) int {
	cycle := 1
	X := 1
	cycles := []int{}
	for idx := 0; idx <= 220; idx++ {
		cycles = append(cycles, -1)
	}
	cycles = append(cycles, 1)

	for _, cmd := range cmds {
		switch c := cmd.(type) {
		case Add:
			if cycle+1 > 220 {
				break
			}
			cycles[cycle+1] = X
			if cycle+2 > 220 {
				break
			}
			X += c.Args[0]
			cycles[cycle+2] = X
			cycle += 2
		case Noop:
			cycle++
			if cycle > 220 {
				break
			}
			cycles[cycle] = X
		}
		if cycle > 220 {
			break
		}

	}
	return 20*cycles[20] + 60*cycles[60] + 100*cycles[100] + 140*cycles[140] + 180*cycles[180] + 220*cycles[220]
}

func main() {
	cmds, err := ReadFromFile("./in-day10.txt")
	if err != nil {
		panic("Could not read from file...")
	}

	fmt.Println(Solve(cmds))
}
