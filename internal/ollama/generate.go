package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func (o *Ollama) Generate(prompt string) MessageResponseType {
	ch := make(chan string, 1)
	var errRet error

	defer func() {
		if errRet != nil {
			ch <- errRet.Error()
			close(ch)
		}
	}()

	url := o.host + "/api/generate"
	requestBody, err := json.Marshal(map[string]string{
		"model":  "llama3",
		"prompt": prompt,
	})
	if err != nil {
		errRet = fmt.Errorf("error creating JSON request body: %s", err.Error())
		return ch
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		errRet = fmt.Errorf("error creating POST request: %s", err.Error())
		return ch
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errRet = fmt.Errorf("error making HTTP request: %s", err.Error())
		return ch
	}

	defer func() {
		if errRet != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		errRet = fmt.Errorf("error, response status code %d", resp.StatusCode)
		return ch
	}

	go func() {
		defer func() {
			resp.Body.Close()
			close(ch)
		}()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()

			var res Response
			if err := json.Unmarshal([]byte(line), &res); err != nil {
				ch <- fmt.Sprintf("Cannot unmarshal: %s", line)
				continue
			}

			ch <- res.Response
		}

		if err := scanner.Err(); err != nil {
			ch <- fmt.Sprintf("error scanning: %s", err.Error())
		}
	}()

	return ch
}
