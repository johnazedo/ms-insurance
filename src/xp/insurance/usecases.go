package insurance

// InsuranceRepository A repository interface to retrieve data of an insurance and cancel it too
type InsuranceRepository interface {
	GetInsuranceInfo(userID string) (*InsuranceInfo, error)
	CancelInsurance(userID string) (*InsuranceInfo, error)
}

type GetInsuranceInfoUseCase struct {
	InsuranceRepository
}

// Invoke Just a proxy usecase to keep the code on the chosen architecture
func (uc *GetInsuranceInfoUseCase) Invoke(userID string) (*InsuranceInfo, error) {
	return uc.GetInsuranceInfo(userID)
}

type CancelInsuranceUseCase struct {
	InsuranceRepository
}

// Invoke Just a proxy usecase to keep the code on the chosen architecture
func (uc *CancelInsuranceUseCase) Invoke(userID string) (*InsuranceInfo, error) {
	return uc.CancelInsurance(userID)
}
