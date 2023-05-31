package cancel

type Request struct {
	UserID string `json:"user-id"`
}

type Response struct {
	Status string `json:"insurance-status"`
}