package ollama

func New(host string) OllamaManager {
	return &Ollama{
		host: host,
		Conversation: &Messages{
			Model:    "llama3",
			Messages: make([]Message, 0),
		},
	}
}

type MessageResponseType = chan string

type OllamaManager interface {
	Generate(string) MessageResponseType
	Chat(string) MessageResponseType
}

type Ollama struct {
	host         string
	Conversation *Messages
}
