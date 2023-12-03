package main

import (
	"fmt"
	"log"
	"strconv"
)

func unscramble(input []byte) int {
	sum := 0
	var firstDigitStr *string
	for _, b := range input {
		_, err := strconv.Atoi(string(b))
		if err != nil {
			continue
		}
		if firstDigitStr == nil {
			firstDigitStr = func() *string {
				s := string(b)
				return &s
			}()
			continue
		}
		twoDigitNum := *firstDigitStr + string(b)
		fmt.Println(twoDigitNum)
		n, err := strconv.Atoi(twoDigitNum)
		if err != nil {
			log.Printf("Error converting string to int: %s", err)
		}
		sum += n
		firstDigitStr = nil
	}
	return sum
}
