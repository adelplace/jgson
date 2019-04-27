package main

import (
	"fmt"
	"os"
)

func main() {
	json := os.Args[1]

	indentLevel := 0
	inText := false
	inValue := false
	ignore := false
	for _, char := range json {
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
				inValue = false
			}
		}

		fmt.Printf("%c", char)

		switch char {
		case '{':
			eol()
			indentLevel++
			indent(indentLevel)
			inValue = false
		case ',':
			eol()
			indent(indentLevel)
		case ':':
			space()
			inValue = true
		case '[':
			eol()
			indentLevel++
			indent(indentLevel)
		case '"':
			if !inText || ignore {
				break
			}

			if inValue {
				beginGreenColor()
			} else {
				beginBlueColor()
			}
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
