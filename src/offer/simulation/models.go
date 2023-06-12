package simulation

type Response struct {
	PhoneBrand     string  `json:"cellphone-brand"`
	PhoneModel     string  `json:"cellphone-model"`
	ValuePerMonth  float64 `json:"value-per-month"`
	Franchise 	   float64 `json:"franchise"`
}

type Request struct {
	UserID 		   string  `json:"user-id" binding:"required"`
	PhoneBrandCode string  `json:"cellphone-brand-code" binding:"required"`
	PhoneModelCode string  `json:"cellphone-model-code" binding:"required"`
}

type BuyResponse struct {
	*Response
	PaymentID string `json:"payment-id"`
}

type BuyErrorResponse struct {
	*Response
	Error string `json:"error-description"`
}

type PhoneInfo struct {
	PhoneBrand     string
	PhoneModel     string
	ValuePerMonth  float64
	Franchise 	   float64
}

type PaymentInfo struct {
	ID string
}