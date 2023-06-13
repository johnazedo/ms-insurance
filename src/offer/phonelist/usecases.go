package phonelist

// This repository will be implemented to retrieve infomation about 
type CellphonesAvailableRepository interface {
	GetListOfBrands() ([]Brand, error)
	GetListOfModels(brandCode string) ([]Model, error)
	GetBrandByCode(brandCode string) (Brand, error)
}

type GetListOfBrandsUseCase struct {
	CellphonesAvailableRepository
}

// Just a proxy usecase to keep the code on the chosen archtecture
func (uc *GetListOfBrandsUseCase) Invoke() ([]Brand, error) {
	return uc.GetListOfBrands()
}

type GetListOfModelsUseCase struct {
	CellphonesAvailableRepository
}

// GetListOfModelsUseCase.Invoke get the list of models for the brand. If the operation
// was successful the function will return information about the Brand too.
func (uc *GetListOfModelsUseCase) Invoke(brandCode string) (*Brand, []Model, error) {
	brand, err := uc.GetBrandByCode(brandCode)
	if err != nil {
		return nil, nil, err
	}

	models, err := uc.GetListOfModels(brand.Code)
	if err != nil {
		return nil, nil, err
	}

	return &brand, models, nil
}
