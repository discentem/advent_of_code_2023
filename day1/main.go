package main

import (
	"fmt"
	"log"
	"strconv"
)

type line string

func (l line) onlyNums() line {
	var final line
	for _, b := range l {
		n, err := strconv.Atoi(string(b))
		if err != nil {
			continue
		}
		final = final + line(strconv.Itoa(n))
	}
	fmt.Println("final: ", final)
	return final
}

func (l line) numFromIndicies(indicies []int) (*int, error) {
	final := ""
	for _, i := range indicies {
		if len(l)-1 < i {
			return nil, fmt.Errorf("invalid index %d", i)
		}
		_, err := strconv.Atoi(final)
		if err != nil {
			continue
		}
		final = final + string(l[i])
	}
	n, err := strconv.Atoi(final)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func split(input []byte) []line {
	var lines []line
	var current line
	for _, b := range input {
		if b == '\n' {
			lines = append(lines, current)
			current = ""
			continue
		}
		current = current + line(string(b))
	}
	if current != "" {
		lines = append(lines, current)
	}
	fmt.Println("lines:", lines)
	return lines
}

func unscramble(input []byte) int {
	var sum int
	for _, l := range split(input) {
		fmt.Println(l)
		l = l.onlyNums()
		fmt.Println(l)
		n, err := l.numFromIndicies([]int{0, len(l) - 1})
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + *n
	}
	return sum
}
