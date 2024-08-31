package speach

import (
	"log"
	"os/exec"
	"strings"
)

func TextToSPeach(text string) {

	go func() {
		cleanedText := strings.ReplaceAll(text, "**", "")
		cmd := exec.Command("espeak", cleanedText)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
