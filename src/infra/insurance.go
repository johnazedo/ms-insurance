package infra

import (
	"strconv"
	"time"

	. "github.com/johnazedo/ms-insurance/src/xp/insurance"
	"gorm.io/gorm"
)

type InsuranceRepositoryImpl struct {
	DB *gorm.DB
}

func (r *InsuranceRepositoryImpl) GetInsuranceInfo(userID string) (*InsuranceInfo, error) {
	var insuranceInfoModel InsuranceInfoModel
	result := r.DB.Where("user_id = ?", userID).First(&insuranceInfoModel)
	if result.Error != nil {
		return nil, result.Error
	}

	var model ModelModel
	result = r.DB.First(&model, insuranceInfoModel.ModelID)
	if result.Error != nil {
		return nil, result.Error
	}

	var brand BrandModel
	result = r.DB.First(&brand, model.BrandID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &InsuranceInfo{
		ID: strconv.FormatUint(uint64(insuranceInfoModel.ID), 10) ,
		PhoneBrand: brand.Name,
		PhoneModel: model.Name,
		ValuePerMonth: model.ValuePerMonth,
		Franchise: model.Franchise,
		Validity:  insuranceInfoModel.Validity,
		Status: insuranceInfoModel.Status,
		UserID: insuranceInfoModel.UserID,
	}, nil
}

func (r *InsuranceRepositoryImpl) CancelInsurance(userID string) (*InsuranceInfo, error) {
	result := r.DB.Model(&InsuranceInfoModel{}).Where("user_id = ?", userID).Update("status", "CANCELLED")
	if result.Error != nil {
		return nil, result.Error
	}
	return r.GetInsuranceInfo(userID)
}

func (r *InsuranceRepositoryImpl) SaveInsurance(userID string, paymentID string, phoneBrand string, phoneModel string) error {
	var model ModelModel
	result := r.DB.Where("code = ?", phoneModel).First(&model)
	if result.Error != nil {
		return result.Error
	}

	insurance := InsuranceInfoModel{
		Validity: time.Now().AddDate(0, 1, 0).String(),
		Status: "ACTIVE",
		ModelID: strconv.FormatUint(uint64(model.ID), 10),
		UserID: userID,
	}

	result = r.DB.Create(&insurance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *InsuranceRepositoryImpl) GetInsurances() ([]*InsuranceInfo, error) {
	var models []ModelModel
	var insurances []*InsuranceInfo
	result := r.DB.Joins("InsuranceInfoModel").Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, model := range models {
		for _, insurance := range model.Insurance {
			insurances = append(insurances, &InsuranceInfo{
				ID: strconv.FormatUint(uint64(insurance.ID), 10),
				PhoneBrand: "",
				PhoneModel: model.Name,
				ValuePerMonth: model.ValuePerMonth,
				Franchise: model.Franchise,
				Validity:  insurance.Validity,
				Status: insurance.Status,
				UserID: insurance.UserID,
			})
		}
	}

	return insurances, nil
}

func (r *InsuranceRepositoryImpl) UpdateValidityInsurance(id string, newValidity string) error {
	result := r.DB.Model(&InsuranceInfoModel{}).Where("id = ?", id).Update("validity", newValidity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
