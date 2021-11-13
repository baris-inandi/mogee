package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

const (
	tapeLen = 32
)

func getCode(f string) string {
	fileBytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}
	return string(fileBytes)
}

func evalExpr(code string, ptr uint, tape [tapeLen]byte) ([tapeLen]byte, uint) {
	skipChars := 0
	for _, char := range code {
		if skipChars > 0 {
			skipChars--
		} else {
			char := string(char)
			_, size := utf8.DecodeRuneInString(char)
			if size > 1 {
				switch char {

				// pointer movement
				case "👈":
					ptr--
				case "👉":
					ptr++
				case "🚘":
					ptr = ptr + 10
				case "🚗":
					ptr = ptr - 10

				// increment/decrement functions
				case "👍":
					tape[ptr]++
				case "👎":
					tape[ptr]--
				case "✋":
					tape[ptr] = tape[ptr] + 5
				case "🤚":
					tape[ptr] = tape[ptr] - 5
				case "🔵":
					tape[ptr] = tape[ptr] + 10
				case "🟦":
					tape[ptr] = tape[ptr] - 10
				case "🔴":
					tape[ptr] = tape[ptr] + 100
				case "🟥":
					tape[ptr] = tape[ptr] - 100

				// i/o
				case "📝":
					fmt.Printf(string(tape[ptr]))
				case "📖":
					var bfIn byte
					fmt.Printf("\n📖 ")
					fmt.Scanln(&bfIn)
					tape[ptr] = bfIn

				// shortcuts
				case "🧿":
					tape[ptr] = 255
				case "🚫":
					tape[ptr] = 0
				case "🚿":
					tape = [tapeLen]byte{}
				case "❌":
					os.Exit(0)
				case "📼":
					fmt.Println(tape)
				}
			}
		}
	}
	return tape, ptr
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Emoji Language Interpreter")
		fmt.Println("Usage: ./main <file>")
		os.Exit(0)
	}
	code := getCode(args[0])
	evalExpr(code, 0, [tapeLen]byte{})
}
