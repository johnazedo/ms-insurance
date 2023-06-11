package simulation

// Simulation usecases
type PhoneInfoRepository interface {
	SearchByBrandAndModel(phoneBrand string, phoneModel string) (*PhoneInfo, error)
}

type GetPhoneInformationUseCase struct {
	repository PhoneInfoRepository
}

func (uc *GetPhoneInformationUseCase) Invoke(phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	return uc.repository.SearchByBrandAndModel(phoneBrand, phoneModel)
}


// Buy inusrance usecases
type PaymentRepository interface {
	MakePayment(userID string, value float64) error
}

type BuyInsuranceRepository interface {
	SaveInsurance(userID string, phoneBrand string, phoneModel string) error 
}

type BuyInsuranceUseCase struct {
	PaymentRepository
	BuyInsuranceRepository
	PhoneInfoRepository
}

func (uc *BuyInsuranceUseCase) Invoke(userID string, phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	info, err := uc.SearchByBrandAndModel(phoneBrand, phoneModel)
	if err != nil {
		return nil, err
	}

	err = uc.MakePayment(userID, info.ValuePerMonth)
	if err != nil {
		return nil, err
	}

	err = uc.SaveInsurance(userID, phoneBrand, phoneModel)
	if err != nil {
		return nil, err
	}

	return info, err
}