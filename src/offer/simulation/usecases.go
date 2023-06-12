package simulation

// Simulation usecases
type PhoneInfoRepository interface {
	SearchByBrandAndModel(phoneBrand string, phoneModel string) (*PhoneInfo, error)
}

type GetPhoneInformationUseCase struct {
	PhoneInfoRepository
}

func (uc *GetPhoneInformationUseCase) Invoke(phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	return uc.SearchByBrandAndModel(phoneBrand, phoneModel)
}


// Buy inusrance usecases
type PaymentRepository interface {
	MakePayment(userID string, value float64) (string, error)
}

type BuyInsuranceRepository interface {
	SaveInsurance(userID string, paymentID string, phoneBrand string, phoneModel string) error 
}

type BuyInsuranceUseCase struct {
	PaymentRepository
	BuyInsuranceRepository
	PhoneInfoRepository
}

func (uc *BuyInsuranceUseCase) Invoke(userID string, phoneBrand string, phoneModel string) (*PhoneInfo, *PaymentInfo, error) {
	info, err := uc.SearchByBrandAndModel(phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	paymentID, err := uc.MakePayment(userID, info.ValuePerMonth)
	if err != nil {
		return nil, nil, err
	}
	paymentInfo := PaymentInfo{
		ID: paymentID,
	}

	err = uc.SaveInsurance(userID, paymentID, phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	return info, &paymentInfo, err
}