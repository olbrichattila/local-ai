package ollama

func New(host string) OllamaManager {
	return &Ollama{host: host}
}

type MessageResponseType = chan string

type OllamaManager interface {
	Generate(string) MessageResponseType
}

type Ollama struct {
	host string
}
