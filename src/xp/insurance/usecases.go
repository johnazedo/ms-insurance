package insurance

// A repository interface to retrieve data of a insurance and cancel it too
type InsuranceRepository interface {
	GetInsuranceInfo(userID string) (*InsuranceInfo, error)
	CancelInsurance(userID string) (*InsuranceInfo, error)
}

type GetInsuranceInfoUseCase struct {
	InsuranceRepository
}

// Just a proxy usecase to keep the code on the chosen archtecture
func (uc *GetInsuranceInfoUseCase) Invoke(userID string) (*InsuranceInfo, error) {
	return uc.GetInsuranceInfo(userID)
}

type CancelInsuranceUseCase struct {
	InsuranceRepository
}

// Just a proxy usecase to keep the code on the chosen archtecture
func (uc *CancelInsuranceUseCase) Invoke(userID string) (*InsuranceInfo, error) {
	return uc.CancelInsurance(userID)
}
