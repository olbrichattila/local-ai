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

			for resp := range generated {
				// fmt.Printf("|%s|", resp)
				// speach.TextToSPeach(resp)
				firstChar := "-"
				if len(resp) > 0 {
					firstChar = string(resp[0])
				}
				fullWord = fullWord + resp
				if firstChar == " " {
					fmt.Print(fullWord)
					speach.TextToSPeach(fullWord)
					fullWord = ""
				}

			}

			if fullWord != "" {
				fmt.Print(fullWord)
				speach.TextToSPeach(fullWord)
			}

			// sentence := ""

			// for resp := range generated {
			// 	fmt.Print(resp)
			// 	// speach.TextToSPeach(resp)
			// 	lastChar := " "
			// 	if len(resp) > 0 {
			// 		lastChar = string(resp[len(resp)-1])
			// 	}

			// 	sentence = sentence + resp
			// 	if lastChar == "." || lastChar == "!" || lastChar == "\n" {
			// 		speach.TextToSPeach(sentence)
			// 		sentence = ""
			// 	}
			// }

			// if sentence != "" {
			// 	speach.TextToSPeach(sentence)
			// }

			fmt.Println()
		} else {
			break
		}
	}
}
