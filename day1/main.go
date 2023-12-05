package main

import (
	"fmt"
	"strconv"
)

type line string

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

func unscramble(input []byte) (int, error) {
	var sum int
	var first *string
	var last string
	for _, l := range split(input) {
		for _, c := range l {
			_, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}
			if first == nil {
				sc := string(c)
				first = &sc
				continue
			}
			last = string(c)
		}
		if last == "" {
			last = *first
		}
		num := *first + last
		n, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		fmt.Println("n:", n)
		sum = sum + n
		first = nil
		last = ""
	}
	return sum, nil
}
