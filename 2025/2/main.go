package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	firstNumber int = iota
	secondNumber
)

func main() {
	text, err := os.ReadFile("./data/input.txt")
	if err != nil {
		panic("wheres is your fileeeeeeeeeeee")
	}

	answer := 0
	var initN, finalN strings.Builder
	mode := firstNumber

	for _, b := range text {
		if b == 10 {
			continue
		}
		switch mode {
		case firstNumber:
			if b == '-' {
				mode = secondNumber
				continue
			}
			initN.WriteByte(b)
		case secondNumber:
			if b == ',' {
				answer += processSequence(initN.String(), finalN.String())
				initN.Reset()
				finalN.Reset()
				mode = firstNumber
				continue
			}

			finalN.WriteByte(b)
		}
	}
	answer = answer + processSequence(initN.String(), finalN.String())
	fmt.Println(answer)
}

func processSequence(initN, finalN string) (invalidsSum int) {
	iN, err := strconv.Atoi(initN)
	if err != nil {
		panic("tf?")
	}

	fN, err := strconv.Atoi(finalN)
	if err != nil {
		panic("tf?")
	}

	for i := iN; i <= fN; i++ {
		invalidsSum += checkInvalid(i)
	}

	return invalidsSum
}

func checkInvalid(x int) int {
	s := strconv.Itoa(x)

	if len(s)%2 != 0 {
		return 0
	}

	half := len(s) / 2

	if s[:half] != s[half:] {
		return 0
	}

	return x
}
