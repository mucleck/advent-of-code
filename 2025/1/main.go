package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	letter int = iota
	number
)

func main() {
	text, err := os.ReadFile("./data/input.txt")
	if err != nil {
		panic("cannot even read bruh")
	}

	var cN strings.Builder
	mode := letter
	dir := false
	pos := 50
	password := 0

	for _, c := range text {
		switch mode {
		case letter:
			if c == 'L' {
				dir = true
			} else {
				dir = false
			}
			mode = number
		case number:
			if c == '\n' {
				n, _ := strconv.Atoi(cN.String())
				if dir {
					n = n * -1
				}
				pos = mod((pos + n), 100)
				if pos == 0 {
					password++
				}
				mode = letter
				cN.Reset()
				continue
			}
			cN.WriteByte(c)
		}
	}
	fmt.Println(password)
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
