package simulation

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// Put usecase here
}

// GetInsuranceSimulation godoc
// @Summary Show simulation data.
// @Description Show simulation data.
// @Schemes
// @Tags Simulation
// @Accept */*
// @Produce json
// @Param body body Request true "Resquest to simulation info"
// @Success 200 {object} Response
// @Router /simulation [post]
func (ctrl *Controller) GetInsuranceSimulation(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := Response {
		UserID: request.UserID,
		PhoneBrandCode: request.PhoneBrandCode,
		PhoneModelCode: request.PhoneModelCode,
		ValuePerMonth: 60.00,
		Franchise: 1200.00,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

// BuyInsurance godoc
// @Summary Buy new insurance.
// @Description Buy new insurance.
// @Tags Simulation
// @Accept */*
// @Param body body Request true "Buy a new insurance"
// @Produce json
// @Success 200 {object} BuyResponse
// @Router /buy [post]
func (ctrl *Controller) BuyInsurance(ctx *gin.Context) {
	var request Request

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := Response {
		UserID: request.UserID,
		PhoneBrandCode: request.PhoneBrandCode,
		PhoneModelCode: request.PhoneModelCode,
		ValuePerMonth: 60.00,
		Franchise: 1200.00,
	}

	buyResponse := BuyResponse {
		Response: response,
		PaymentID: "uuasdjf-aidfnkd-adsfksn",
	}

	ctx.IndentedJSON(http.StatusOK, buyResponse)
}