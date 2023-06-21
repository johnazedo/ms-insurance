package simulation

// Simulation usecases

// PhoneInfoRepository This repository is necessary to get cellphone info using the brand and model
// chosen by client
type PhoneInfoRepository interface {
	SearchByBrandAndModel(phoneBrand string, phoneModel string) (*PhoneInfo, error)
}

type GetPhoneInformationUseCase struct {
	phoneInfoRepository PhoneInfoRepository
}

// Invoke Just a proxy usecase to keep the code on the chosen architecture
func (uc *GetPhoneInformationUseCase) Invoke(phoneBrand string, phoneModel string) (*PhoneInfo, error) {
	return uc.phoneInfoRepository.SearchByBrandAndModel(phoneBrand, phoneModel)
}

// Buy insurance use cases

// PaymentRepository This repository will be implemented to communicate with payment microservice
type PaymentRepository interface {
	MakePayment(userID string, value float64) (*PaymentInfo, error)
}

// BuyInsuranceRepository This repository will be used to store data about insurance on database
type BuyInsuranceRepository interface {
	SaveInsurance(userID string, paymentID string, phoneBrand string, phoneModel string) error
}

type BuyInsuranceUseCase struct {
	paymentRepository      PaymentRepository
	buyInsuranceRepository BuyInsuranceRepository
	phoneInfoRepository    PhoneInfoRepository
}

// Invoke is a method to buy a cellphone insurance and return the information of cellphone and payment.
// On this function we need to get information of phone with the brand and model by code and use the value per month to
// generate a payment. Finally, this function saves the information about new insurance contract on the database
func (uc *BuyInsuranceUseCase) Invoke(userID string, phoneBrand string, phoneModel string) (*PhoneInfo, *PaymentInfo, error) {
	info, err := uc.phoneInfoRepository.SearchByBrandAndModel(phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	paymentInfo, err := uc.paymentRepository.MakePayment(userID, info.ValuePerMonth)
	if err != nil {
		return nil, nil, err
	}

	err = uc.buyInsuranceRepository.SaveInsurance(userID, paymentInfo.ID, phoneBrand, phoneModel)
	if err != nil {
		return nil, nil, err
	}

	return info, paymentInfo, err
}
