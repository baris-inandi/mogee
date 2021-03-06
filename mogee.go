package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	tapeLen = 10000
)

var funcStore = make(map[string]string)
var ptr uint = 0
var tape = [tapeLen]byte{}

func getCode(f string) string {
	fileBytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}
	return string(fileBytes)
}

func trackFunc(code string, index int) int {
	code = code[index+4:]
	code = strings.Split(code, "š")[0]
	fID := code[:4]
	fBody := code[4:]
	funcStore[fID] = fBody
	return len([]rune(code))
}

func evalExpr(code string) {
	skipChars := 0
	callState := false
	for index, char := range code {
		if skipChars > 0 {
			skipChars--
		} else {
			char := string(char)
			_, size := utf8.DecodeRuneInString(char)
			if size > 1 {
				switch char {

				// pointer control
				case "š":
					ptr--
				case "š":
					ptr++
				case "š":
					ptr = ptr + 10
				case "š":
					ptr = ptr - 10

				// increment/decrement functions
				case "š":
					tape[ptr]++
				case "š":
					tape[ptr]--
				case "ā":
					tape[ptr] = tape[ptr] + 5
				case "š¤":
					tape[ptr] = tape[ptr] - 5
				case "šµ":
					tape[ptr] = tape[ptr] + 10
				case "š¦":
					tape[ptr] = tape[ptr] - 10
				case "š“":
					tape[ptr] = tape[ptr] + 100
				case "š„":
					tape[ptr] = tape[ptr] - 100

				// i/o
				case "š":
					fmt.Print(string(tape[ptr]))
				case "š":
					var bfIn byte
					fmt.Print("\nš ")
					fmt.Scanln(&bfIn)
					tape[ptr] = bfIn
				case "š§®":
					fmt.Print(tape[ptr])

				// shortcuts
				case "š§æ":
					tape[ptr] = 255
				case "š«":
					tape[ptr] = 0
				case "šæ":
					tape = [tapeLen]byte{}
				case "ā":
					os.Exit(0)
				case "š¼":
					fmt.Println(tape)

				// functions
				case "š":
					skip := trackFunc(code, index)
					skipChars = skip
				case "š":
					callState = true
				default:
					if callState {
						callState = false
						evalExpr(funcStore[char])
					}
				}
			}
		}
	}
}

func readFromFile() string {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("mogee Interpreter")
		fmt.Println("Usage: ./mogee <file>")
		os.Exit(0)
	}
	return getCode(args[0])
}

func main() {
	evalExpr(readFromFile())
}
