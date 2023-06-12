package phonelist

type CellphonesAvailableRepositoryImpl struct {
	CellphonesAvailableRepository
}

func (r *CellphonesAvailableRepositoryImpl) GetListOfBrands() ([]Brand, error) {
	brands := []Brand{
		{
			Code: "0001",
			Name: "Sansumg",
		},
		{
			Code: "0002",
			Name: "Motorola",
		},
		{
			Code: "0003",
			Name: "Apple",
		},
	}

	return brands, nil
}

func (r *CellphonesAvailableRepositoryImpl) GetListOfModels(brandCode string) ([]Model, error) {
	models := []Model{
		{
			Code: "0001",
			Name: "S22 128GB",
		},
		{
			Code: "0002",
			Name: "S22 256GB",
		},
		{
			Code: "0003",
			Name: "S22 Plus 256GB",
		},
		{
			Code: "0004",
			Name: "S22 Plus 512GB",
		},
	}

	return models, nil
}

func (r *CellphonesAvailableRepositoryImpl) GetBrandByCode(brandCode string) (Brand, error) {
	return Brand{
		Code: "0001",
		Name: "Sansumg",
	}, nil
}
