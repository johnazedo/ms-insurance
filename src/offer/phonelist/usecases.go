package phonelist

type CellphonesAvailableRepository interface {
	GetListOfBrands() ([]Brand, error)
	GetListOfModels(brandCode string) ([]Model, error)
	GetBrandByCode(brandCode string) (Brand, error)
}

type GetListOfBrandsUseCase struct {
	CellphonesAvailableRepository
}

func (uc *GetListOfBrandsUseCase) Invoke() ([]Brand, error) {
	return uc.GetListOfBrands()
}

type GetListOfModelsUseCase struct {
	CellphonesAvailableRepository
}

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
