package main

import (
	"bufio"
	"chat/internal/ollama"
	"chat/internal/speach"
	"fmt"
	"os"
)

func main() {
	lm := ollama.New("http://localhost:11434")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred:", err)
			return
		}

		input = input[:len(input)-1]
		if len(input) > 0 {
			// generated := lm.Generate(input)
			generated := lm.Chat(input)
			fullWord := ""
			speach.Speaker()

			for resp := range generated {
				firstChar := "-"
				if len(resp) > 0 {
					firstChar = string(resp[0])
				}
				fullWord = fullWord + resp
				if firstChar == " " {
					fmt.Print(fullWord)
					speach.Append(fullWord)
					fullWord = ""
				}

			}

			if fullWord != "" {
				speach.Append(fullWord)
				fmt.Print(fullWord)
			}

			fmt.Println()
		} else {
			break
		}
	}
}
