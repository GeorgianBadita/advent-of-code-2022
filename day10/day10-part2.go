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

func Solve(cmds []Command) {
	cycle := 1
	X := 1
	cycles := []int{}
	cycles = append(cycles, 1)

	for _, cmd := range cmds {
		switch c := cmd.(type) {
		case Add:
			cycles = append(cycles, X)
			X += c.Args[0]
			cycles = append(cycles, X)
			cycle += 2
		case Noop:
			cycle++
			cycles = append(cycles, X)
		}
	}

	string := ""
	for idx := 0; idx < 240; idx++ {
		spirte := cycles[idx]
		if idx%40 >= spirte-1 && idx%40 <= spirte+1 {
			string += "█"
		} else {
			string += "."
		}
		if (idx+1)%40 == 0 {
			string += "\n"
		}
	}

	fmt.Println(string)
}

func main() {
	cmds, err := ReadFromFile("./in-day10.txt")
	if err != nil {
		panic("Could not read from file...")
	}

	Solve(cmds)
}
