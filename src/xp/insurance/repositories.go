package insurance

type InsuranceRepositoryImpl struct{}

func (r *InsuranceRepositoryImpl) GetInsuranceInfo(userID string) (*InsuranceInfo, error) {
	return &InsuranceInfo{
		ID:            "00001",
		PhoneBrand:    "Samsung",
		PhoneModel:    "S22 128",
		ValuePerMonth: 60.00,
		Franchise:     1200.00,
		Validity:      "11/05/2024",
		Status:        "ACTIVE",
		UserID:        "00001",
	}, nil
}

func (r *InsuranceRepositoryImpl) CancelInsurance(userID string) (*InsuranceInfo, error) {
	return &InsuranceInfo{
		ID:            "00002",
		PhoneBrand:    "Samsung",
		PhoneModel:    "S22 128",
		ValuePerMonth: 60.00,
		Franchise:     1200.00,
		Validity:      "11/05/2024",
		Status:        "CANCELLED",
		UserID:        "00002",
	}, nil
}