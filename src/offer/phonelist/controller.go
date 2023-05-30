package phonelist

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Controller struct {
	// Put usecase here
}

// GetPhoneBrands godoc
// @Summary Get phone brands list.
// @Description Get phone brands list.
// @Tags PhoneList
// @Accept */*
// @Produce json
// @Success 200 {array} Brand
// @Router /brands [get]
func (ctrl *Controller) GetPhoneBrands(ctx *gin.Context) {

	response := []Brand{
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

	ctx.IndentedJSON(http.StatusOK, response)
}

// GetPhoneModels godoc
// @Summary Get phone models list.
// @Description Get phone models list.
// @Tags PhoneList	
// @Accept */*
// @Produce json
// @Param cellphone-brand-code query string true "A brand code of cellphone"
// @Success 200 {object} ModelResponse
// @Router /models [get]
func (ctrl *Controller) GetPhoneModels(ctx *gin.Context) {
	param, has := ctx.GetQuery("cellphone-brand-code")
	if !has {
		ctx.JSON(http.StatusBadRequest, "message: Brand code was not sent")
		return
	}

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

	response := ModelResponse {
		BrandCode: param,
		BrandName: "Sansumg",
		ModelList: models,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}