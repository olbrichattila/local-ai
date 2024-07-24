package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Messages struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

func (o *Ollama) Chat(prompt string) MessageResponseType {
	o.Conversation.Messages = append(o.Conversation.Messages, Message{
		Role:    "user",
		Content: prompt,
	})

	ch := make(chan string, 1)
	var errRet error

	defer func() {
		if errRet != nil {
			ch <- errRet.Error()
			close(ch)
		}
	}()

	url := o.host + "/api/chat"

	requestBody, err := json.Marshal(o.Conversation)
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

		fullResponse := strings.Builder{}

		for scanner.Scan() {
			line := scanner.Text()

			var res ChatResponse
			if err := json.Unmarshal([]byte(line), &res); err != nil {
				ch <- fmt.Sprintf("Cannot unmarshal: %s", line)
				continue
			}

			if res.Done {
				o.Conversation.Messages = append(o.Conversation.Messages, Message{
					Role:    res.Message.Role,
					Content: fullResponse.String(),
				})
			} else {
				fullResponse.Write([]byte(res.Message.Content))
				ch <- res.Message.Content
			}
		}

		if err := scanner.Err(); err != nil {
			ch <- fmt.Sprintf("error scanning: %s", err.Error())
		}
	}()

	return ch
}
