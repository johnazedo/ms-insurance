package insurance

func MapperFromInsuranceInfoToResponse(info *InsuranceInfo) *Response {
	return &Response{
		PhoneBrand:    info.PhoneBrand,
		PhoneModel:    info.PhoneModel,
		ValuePerMonth: info.ValuePerMonth,
		Franchise:     info.Franchise,
		Validity:      info.Validity,
		Status:        info.Status,
	}
}