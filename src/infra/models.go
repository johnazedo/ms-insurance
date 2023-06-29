package infra

import "gorm.io/gorm"

type InsuranceInfoModel struct {
	gorm.Model
	Validity string
	Status   string
	ModelID  string
	UserID   string
}

type BrandModel struct {
	gorm.Model
	Code       string
	Name       string
	ModelModel []ModelModel `gorm:"foreignKey:BrandID"`
}

type ModelModel struct {
	gorm.Model
	Code          string
	Name          string
	ValuePerMonth float64
	Franchise     float64
	BrandID       string
	Insurance     []InsuranceInfoModel `gorm:"foreignKey:ModelID"`
}

