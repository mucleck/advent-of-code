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

	var steps []string
	var cN strings.Builder
	mode := letter

	for _, c := range text {
		switch mode {
		case letter:
			steps = append(steps, string(c))
			mode = number
		case number:
			if c == '\n' {
				mode = letter
				steps = append(steps, cN.String())
				cN.Reset()
				continue
			}

			cN.WriteByte(c)
		}
	}

	password := 0
	pos := 50
	mode = letter
	var direction bool
	for _, w := range steps {
		if mode == letter {
			if w == "L" {
				direction = true
			} else {
				direction = false
			}
			mode = number
		} else {
			n, _ := strconv.Atoi(w)
			if direction {
				n = n * -1
			}
			pos = mod((pos + n), 100)
			if pos == 0 {
				password++
			}
			mode = letter
		}
	}

	fmt.Println(password)
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
