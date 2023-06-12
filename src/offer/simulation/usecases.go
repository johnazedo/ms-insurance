package simulation

// Simulation usecases
type PhoneInfoRepository interface {
	SearchByBrandAndModel(phoneBrand string, phoneModel string) (*PhoneInfo, error)
}

type GetPhoneInformationUseCase struct {
	PhoneInfoRepository
}

// Just a proxy usecase to keep the code on the chosen archtecture
func (uc *GetPhoneInformationUseCase) Invoke(phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	return uc.SearchByBrandAndModel(phoneBrand, phoneModel)
}

// Buy inusrance usecases
type PaymentRepository interface {
	MakePayment(userID string, value float64) (*PaymentInfo, error)
}

type BuyInsuranceRepository interface {
	SaveInsurance(userID string, paymentID string, phoneBrand string, phoneModel string) error
}

type BuyInsuranceUseCase struct {
	PaymentRepository
	BuyInsuranceRepository
	PhoneInfoRepository
}

// BuyInsuranceUseCase.Invoke is a method to buy a cellphone insurance and return the infomation of cellphone and payment.
// On this function we need to get information of phone with the brand and model by code and use the value per month to
// generate a payment. Finally this function saves the information about new insurance contract on the database.
func (uc *BuyInsuranceUseCase) Invoke(userID string, phoneBrand string, phoneModel string) (*PhoneInfo, *PaymentInfo, error) {
	info, err := uc.SearchByBrandAndModel(phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	paymentInfo, err := uc.MakePayment(userID, info.ValuePerMonth)
	if err != nil {
		return nil, nil, err
	}

	err = uc.SaveInsurance(userID, paymentInfo.ID, phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	return info, paymentInfo, err
}
