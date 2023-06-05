package simulation

// Simulation usecases
type PhoneInfoRepository interface {
	SearchByBrandAndModel(phoneBrand string, phoneModel string) (PhoneInfo, error)
}

type GetPhoneInformationUseCase struct {
	repository PhoneInfoRepository
}

func (uc *GetPhoneInformationUseCase) Invoke(phoneBrand string, phoneModel string) (PhoneInfo, error) {
	return uc.repository.SearchByBrandAndModel(phoneBrand, phoneModel)
}


// Buy inusrance usecases
type BuyInsuranceRepository interface {

}

type BuyInsuranceUseCase struct {

}

func (uc *BuyInsuranceUseCase) Invoke() {
	
}