package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var json []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		json = append(json, input)
	}

	indentLevel := 0
	inText := false
	ignore := false
	for j := 0; j < len(json); j++ {
		char := json[j]
		switch char {
		case ':':
			space()
		case '}':
			eol()
			indentLevel--
			indent(indentLevel)
		case ']':
			eol()
			indentLevel--
			indent(indentLevel)
		case ' ', '\n', '\t':
			if !inText {
				continue
			}
		case '"':
			if ignore {
				break
			}

			inText = !inText
			if !inText {
				endColor()
			}
		}

		fmt.Printf("%c", char)

		switch char {
		case '{':
			eol()
			indentLevel++
			indent(indentLevel)
		case ',':
			eol()
			indent(indentLevel)
		case ':':
			space()
		case '[':
			eol()
			indentLevel++
			indent(indentLevel)
		case '"':
			if !inText || ignore {
				break
			}

			beginGreenColor()
		}

		if char == '\\' {
			ignore = true
		} else {
			ignore = false
		}
	}
}

func eol() {
	fmt.Print("\n")
}

func space() {
	fmt.Print(" ")
}

func indent(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
}

func endColor() {
	fmt.Print("\033[0m")
}

func beginBlueColor() {
	fmt.Print("\033[94m")
}

func beginGreenColor() {
	fmt.Print("\033[92m")
}
