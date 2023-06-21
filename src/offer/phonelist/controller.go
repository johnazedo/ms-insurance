package phonelist

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPhoneListController() Controller {
	repository := &CellphonesAvailableRepositoryImpl{}
	return Controller{
		getListOfBrandsUseCase: GetListOfBrandsUseCase{
			cellphonesAvailableRepository: repository,
		},
		getListOfModelsUseCase: GetListOfModelsUseCase{
			cellphonesAvailableRepository: repository,
		},
	}
}

type Controller struct {
	getListOfBrandsUseCase GetListOfBrandsUseCase
	getListOfModelsUseCase GetListOfModelsUseCase
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
	response, err := ctrl.getListOfBrandsUseCase.Invoke()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
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

	brand, models, err := ctrl.getListOfModelsUseCase.Invoke(param)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ModelResponse{
		BrandCode: brand.Name,
		BrandName: brand.Code,
		ModelList: models,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
