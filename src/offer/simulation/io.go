package simulation

type Response struct {
	PhoneBrand     string  `json:"cellphone-brand"`
	PhoneModel     string  `json:"cellphone-model"`
	ValuePerMonth  float64 `json:"value-per-month"`
	Franchise 	   float64 `json:"franchise"`
}

func MapperFromPhoneInfoToResponse(pi PhoneInfo) *Response {
	return &Response {
		PhoneBrand: pi.PhoneBrand,
		PhoneModel: pi.PhoneModel,
		ValuePerMonth: pi.ValuePerMonth,
		Franchise: pi.Franchise,
	}
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