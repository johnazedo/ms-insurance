package simulation

func MapperFromPhoneInfoToResponse(pi *PhoneInfo) *Response {
	return &Response {
		PhoneBrand: pi.PhoneBrand,
		PhoneModel: pi.PhoneModel,
		ValuePerMonth: pi.ValuePerMonth,
		Franchise: pi.Franchise,
	}
}