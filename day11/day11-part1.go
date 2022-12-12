package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Id         int
	Items      []int
	Div        int
	OnTrue     int
	OnFalse    int
	Multiplier int
	Adder      int
}

func (m Monkey) EvalTest(worry int) bool {
	return worry%m.Div == 0
}

func (m Monkey) ThrowTo(worry int) int {
	if m.EvalTest(worry) {
		return m.OnTrue
	}
	return m.OnFalse
}

func (m Monkey) FinalWorry(worry int) int {
	return (worry*m.Multiplier + m.Adder) / 3
}

func ReadFromFile(filePath string) ([]Monkey, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	res := []Monkey{}

	for scanner.Scan() {
		monkeyNo := strings.Trim(scanner.Text(), " ")
		monkeyNoInt, _ := strconv.Atoi(string(monkeyNo[len(monkeyNo)-1]))

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
		if strings.Contains(operationStr, "+") {
			operationStr := strings.Split(operationStr, "+")
			opInt, _ := strconv.Atoi(operationStr[len(operationStr)-1])
			adder = opInt
		} else {
			operationStr := strings.Split(operationStr, "*")
			opInt, _ := strconv.Atoi(operationStr[len(operationStr)-1])
			multiplier = opInt
		}

		res = append(res, Monkey{Id: monkeyNoInt, Items: items, Multiplier: multiplier, Adder: adder})
	}

	return res, nil
}
