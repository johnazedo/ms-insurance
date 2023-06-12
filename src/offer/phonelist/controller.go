package phonelist

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPhoneListController() controller {
	repository := &CellphonesAvailableRepositoryImpl{}
	return controller{
		GetListOfBrandsUseCase: GetListOfBrandsUseCase{
			CellphonesAvailableRepository: repository,
		},
		GetListOfModelsUseCase: GetListOfModelsUseCase{
			CellphonesAvailableRepository: repository,
		},
	}
}

type controller struct {
	GetListOfBrandsUseCase
	GetListOfModelsUseCase
}

// GetPhoneBrands godoc
// @Summary Get phone brands list.
// @Description Get phone brands list.
// @Tags PhoneList
// @Accept */*
// @Produce json
// @Success 200 {array} Brand
// @Router /brands [get]
func (ctrl *controller) GetPhoneBrands(ctx *gin.Context) {
	response, err := ctrl.GetListOfBrandsUseCase.Invoke()
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
func (ctrl *controller) GetPhoneModels(ctx *gin.Context) {
	param, has := ctx.GetQuery("cellphone-brand-code")
	if !has {
		ctx.JSON(http.StatusBadRequest, "message: Brand code was not sent")
		return
	}

	brand, models, err := ctrl.GetListOfModelsUseCase.Invoke(param)
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
