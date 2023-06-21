package phonelist

// CellphonesAvailableRepository This repository will be implemented to retrieve information about
type CellphonesAvailableRepository interface {
	GetListOfBrands() ([]Brand, error)
	GetListOfModels(brandCode string) ([]Model, error)
	GetBrandByCode(brandCode string) (Brand, error)
}

type GetListOfBrandsUseCase struct {
	cellphonesAvailableRepository CellphonesAvailableRepository
}

// Invoke Just a proxy usecase to keep the code on the chosen architecture
func (uc *GetListOfBrandsUseCase) Invoke() ([]Brand, error) {
	return uc.cellphonesAvailableRepository.GetListOfBrands()
}

type GetListOfModelsUseCase struct {
	cellphonesAvailableRepository CellphonesAvailableRepository
}

// Invoke get the list of models for the brand. If the operation
// was successful the function will return information about the Brand too.
func (uc *GetListOfModelsUseCase) Invoke(brandCode string) (*Brand, []Model, error) {
	brand, err := uc.cellphonesAvailableRepository.GetBrandByCode(brandCode)
	if err != nil {
		return nil, nil, err
	}

	models, err := uc.cellphonesAvailableRepository.GetListOfModels(brand.Code)
	if err != nil {
		return nil, nil, err
	}

	return &brand, models, nil
}
