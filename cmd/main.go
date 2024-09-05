package main

import (
	"bufio"
	"chat/internal/ollama"
	"chat/internal/speach"
	"fmt"
	"os"
	"time"
)

func main() {
	lm := ollama.New("http://localhost:11434")

	// Create a channel to signal the Goroutine to stop
	done := make(chan bool)

	speach.Speaker(done)
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

	done <- true

	time.Sleep(2 * time.Second)
}
