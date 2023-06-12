package insurance

type Request struct {
	UserID string `json:"user-id"`
}

type Response struct {
	PhoneBrand    string  `json:"cellphone-brand"`
	PhoneModel    string  `json:"cellphone-model"`
	ValuePerMonth float64 `json:"value-per-month"`
	Franchise     float64 `json:"franchise"`
	Validity      string  `json:"validity"`
	Status        string  `json:"status"`
}

type InsuranceInfo struct {
	PhoneBrand    string
	PhoneModel    string
	ValuePerMonth float64
	Franchise     float64
	Validity      string
	Status        string
}
