package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Id         int
	Items      []int
	Multiplier int
	Adder      int
	Div        int
	OnTrue     int
	OnFalse    int
	IsSquare   bool
}

func (m Monkey) EvalTest(worry int) bool {
	return worry%m.Div == 0
}

func (m Monkey) FinalWorry(worry int) int {
	mult := m.Multiplier
	if m.IsSquare {
		mult = worry
	}
	return (worry*mult + m.Adder)
}

func ReadFromFile(filePath string) ([]Monkey, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	res := []Monkey{}

	for scanner.Scan() {
		monkeyNo := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		monkeyNoStr := monkeyNo[len(monkeyNo)-1]
		monkeyNoStr = monkeyNoStr[:len(monkeyNoStr)-1]
		monkeyNoInt, _ := strconv.Atoi(monkeyNoStr)

		scanner.Scan()
		itemsLine := strings.Split(strings.Trim(scanner.Text(), " "), ":")
		itemsLine = strings.Split(itemsLine[len(itemsLine)-1], ",")

		items := []int{}
		for _, item := range itemsLine {
			itemInt, _ := strconv.Atoi(strings.Trim(item, " "))
			items = append(items, itemInt)
		}

		scanner.Scan()
		operation := strings.Split(scanner.Text(), "=")
		operationStr := strings.Trim(operation[len(operation)-1], " ")

		multiplier := 1
		adder := 0
		isSq := false

		if strings.Contains(operationStr, "+") {
			operationStr := strings.Split(operationStr, "+")
			opInt, _ := strconv.Atoi(strings.Trim(operationStr[len(operationStr)-1], " "))
			adder = opInt
		} else {
			operationStr := strings.Split(operationStr, "*")
			if strings.Contains(operationStr[len(operationStr)-1], "old") {
				isSq = true
			}
			opInt, _ := strconv.Atoi(strings.Trim(operationStr[len(operationStr)-1], " "))
			multiplier = opInt
		}

		scanner.Scan()
		div := strings.Split(scanner.Text(), "by")
		divStr := strings.Trim(div[len(div)-1], " ")

		divNum, _ := strconv.Atoi(divStr)

		scanner.Scan()
		trueM := strings.Split(scanner.Text(), " ")
		trueMStr := strings.Trim(trueM[len(trueM)-1], " ")

		onTrueMonkey, _ := strconv.Atoi(trueMStr)

		scanner.Scan()
		falseM := strings.Split(scanner.Text(), " ")
		falseMStr := strings.Trim(falseM[len(falseM)-1], " ")

		onFalseMonkey, _ := strconv.Atoi(falseMStr)

		scanner.Scan()

		res = append(res, Monkey{Id: monkeyNoInt, Items: items, Multiplier: multiplier, Adder: adder, Div: divNum, OnTrue: onTrueMonkey, OnFalse: onFalseMonkey, IsSquare: isSq})
	}

	return res, nil
}

func Solve(monkeys []Monkey) int {
	inspected := []int{}
	prod := 1
	for _, monkey := range monkeys {
		inspected = append(inspected, 0)
		prod *= monkey.Div
	}

	fmt.Printf("%v\n", prod)

	for idx := 0; idx < 10000; idx++ {
		for jdx := 0; jdx < len(monkeys); jdx++ {
			for _, worry := range monkeys[jdx].Items {
				newWorry := monkeys[jdx].FinalWorry(worry) % prod
				inspected[jdx]++
				if monkeys[jdx].EvalTest(newWorry) {
					monkeys[monkeys[jdx].OnTrue].Items = append(monkeys[monkeys[jdx].OnTrue].Items, newWorry)
				} else {
					monkeys[monkeys[jdx].OnFalse].Items = append(monkeys[monkeys[jdx].OnFalse].Items, newWorry)
				}
			}
			monkeys[jdx].Items = []int{}
		}
	}
	var max1, max2 int
	if inspected[0] >= inspected[1] {
		max1 = inspected[0]
		max2 = inspected[1]
	} else {
		max2 = inspected[0]
		max1 = inspected[1]
	}

	for idx := 2; idx < len(inspected); idx++ {
		if inspected[idx] >= max1 {
			max2 = max1
			max1 = inspected[idx]
		} else if inspected[idx] > max2 {
			max2 = inspected[idx]
		}
	}

	fmt.Printf("Arr: %v\n", inspected)
	return max1 * max2
}

func main() {
	monkeys, err := ReadFromFile("./in-day11.txt")
	if err != nil {
		panic("Could not read from file...")
	}

	for _, monkey := range monkeys {
		fmt.Printf("%v\n\n\n", monkey)
	}

	fmt.Print(Solve(monkeys))
}
