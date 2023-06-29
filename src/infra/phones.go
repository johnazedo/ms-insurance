package infra

import (
	"gorm.io/gorm"
	phonelist "github.com/johnazedo/ms-insurance/src/offer/phonelist"
	simulation "github.com/johnazedo/ms-insurance/src/offer/simulation"
)

type CellphonesAvailableRepositoryImpl struct {
	DB *gorm.DB
}

func (r *CellphonesAvailableRepositoryImpl) GetListOfBrands() ([]phonelist.Brand, error) {
	var models []BrandModel
	var brands = []phonelist.Brand{}
	result := r.DB.Find(&models)
	if result.Error != nil {
		return brands, result.Error
	}

	for _, brand := range models {
		brands = append(brands, phonelist.Brand{
			Code: brand.Code,
			Name: brand.Name,
		})
	}

	return brands, nil
}

func (r *CellphonesAvailableRepositoryImpl) GetListOfModels(brandCode string) ([]phonelist.Model, error) {
	var dbmodels []ModelModel
	var models = []phonelist.Model{}
	var brand BrandModel
	result := r.DB.Where("code = ?", brandCode).First(&brand)
	if result.Error != nil {
		return models, result.Error
	}

	result = r.DB.Where("brand_id = ?", brand.ID).Find(&dbmodels)
	if result.Error != nil {
		return models, result.Error
	}

	for _, model := range dbmodels {
		models = append(models, phonelist.Model{
			Code: model.Code,
			Name: model.Name,
		})
	}

	return models, nil
}

func (r *CellphonesAvailableRepositoryImpl) GetBrandByCode(brandCode string) (phonelist.Brand, error) {
	var model BrandModel
	var brand phonelist.Brand
	result := r.DB.Where("code = ?", brandCode).First(&model)
	if result.Error != nil {
		return brand, result.Error
	}

	return phonelist.Brand{
		Code: model.Code,
		Name: model.Name,
	}, nil
}

func (r *CellphonesAvailableRepositoryImpl) SearchByBrandAndModel(phoneBrand string, phoneModel string) (*simulation.PhoneInfo, error) {
	var brand BrandModel
	var model ModelModel
	result := r.DB.Where("code = ?", phoneBrand).First(&brand)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.Where("code = ? AND brand_id = ?", phoneModel, brand.ID).First(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return &simulation.PhoneInfo{
		PhoneBrand: brand.Name,
		PhoneModel: model.Name,
		ValuePerMonth: model.ValuePerMonth,
		Franchise: model.Franchise,
	}, nil
}
