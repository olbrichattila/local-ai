package main

import (
	"bufio"
	"chat/internal/ollama"
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

			for resp := range generated {
				fmt.Print(resp)
			}

			fmt.Println()
		} else {
			break
		}
	}
}
