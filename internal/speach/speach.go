package speach

import (
	"log"
	"os/exec"
	"strings"
	"sync"
)

type wordBuffer struct {
	mu    sync.Mutex
	words []string
}

var buffer = &wordBuffer{}

func Append(word string) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()

	buffer.words = append(buffer.words, word)
}

func PullFirst() (string, bool) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()

	if len(buffer.words) == 0 {
		return "", false
	}

	first := buffer.words[0]
	buffer.words = buffer.words[1:]

	return first, true
}

func Speaker(done chan bool) {
	go func() {
		cnt := 1
		fewWords := ""
		for {
			select {
			case <-done:
				return
			default:
				word, ok := PullFirst()
				if ok {
					fewWords = fewWords + word
					cnt--
					if cnt == 0 {
						cnt = 8
						cleanedText := strings.ReplaceAll(fewWords, "**", "")
						fewWords = ""
						cmd := exec.Command("espeak", cleanedText)
						err := cmd.Run()
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}()
}

// func TextToSPeach(text string) {

// 	cleanedText := strings.ReplaceAll(text, "**", "")
// 	cmd := exec.Command("espeak", cleanedText)
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
