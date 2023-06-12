package insurance

type InsuranceRepositoryImpl struct{}

func (r *InsuranceRepositoryImpl) GetInsuranceInfo(userID string) (*InsuranceInfo, error) {
	return &InsuranceInfo{
		PhoneBrand:    "Samsung",
		PhoneModel:    "S22 128",
		ValuePerMonth: 60.00,
		Franchise:     1200.00,
		Validity:      "11/05/2024",
		Status:        "ACTIVE",
	}, nil
}


func (r *InsuranceRepositoryImpl) CancelInsurance(userID string) (*InsuranceInfo, error) {
	return &InsuranceInfo{
		PhoneBrand:    "Samsung",
		PhoneModel:    "S22 128",
		ValuePerMonth: 60.00,
		Franchise:     1200.00,
		Validity:      "11/05/2024",
		Status:        "CANCELLED",
	}, nil
}