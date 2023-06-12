package main

import (
	"github.com/gin-gonic/gin"
	simulation "github.com/johnazedo/ms-insurance/src/offer/simulation"
	phonelist "github.com/johnazedo/ms-insurance/src/offer/phonelist"
	cancel "github.com/johnazedo/ms-insurance/src/xp/cancel"
	getinsurance "github.com/johnazedo/ms-insurance/src/xp/getinsurance"
	swaggerFiles "github.com/swaggo/files"
   	swagger "github.com/swaggo/gin-swagger"
	_ "github.com/johnazedo/ms-insurance/docs"
)

// @title ms-insurance
// @version 1.0
// @description MS with insurance features
// @termsOfService http://swagger.io/terms/

// @contact.name João Pedro Limão
// @contact.email joao.limao.701@ufrn.edu.br

// @securityDefinitions.apiKey JWT
// @in header
// @name token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8001
// @BasePath /

// @schemes http
func main() {
	simulationController := simulation.GetSimulationController()
	phonelistController := phonelist.GetPhoneListController()
	cancelController := cancel.Controller{}
	getinsuranceController := getinsurance.GetInsuranceController()

	router := gin.Default()
	router.POST("/simulation", simulationController.GetInsuranceSimulation)
	router.POST("/buy", simulationController.BuyInsurance)
	router.GET("/brands", phonelistController.GetPhoneBrands)
	router.GET("/models", phonelistController.GetPhoneModels)

	router.POST("/cancel", cancelController.CancelInsurance)
	router.POST("/insurance", getinsuranceController.GetInsuranceInformation)

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	
	router.Run("127.0.0.1:8001")
}