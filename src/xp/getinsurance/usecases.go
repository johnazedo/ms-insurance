package getinsurance

type InsuranceRepository interface {
	GetInsuranceInfo(userID string) (*InsuranceInfo, error)
}

type GetInsuranceInfoUseCase struct {
	InsuranceRepository
}

// Just a proxy usecase to keep the code on the chosen archtecture
func (uc *GetInsuranceInfoUseCase) Invoke(userID string) (*InsuranceInfo, error) {
	return uc.GetInsuranceInfo(userID)
}
