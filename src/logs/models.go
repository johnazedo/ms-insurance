package logs

type Log struct {
	TimeStamp    string `json:"timestamp"`
	Level        string `json:"level"`
	Microservice string `json:"microservice"`
	Thread       string `json:"thread"`
	Class        string `json:"class"`
	Method       string `json:"method"`
	Message      string `json:"message"`
	Context      string `json:"context"`
	Ip           string `json:"ip"`
}

type LogInput struct {
	Level        string
	Class        string
	Method       string
	Message      string
}