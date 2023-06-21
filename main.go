package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/johnazedo/ms-insurance/docs"
	phonelist "github.com/johnazedo/ms-insurance/src/offer/phonelist"
	simulation "github.com/johnazedo/ms-insurance/src/offer/simulation"
	insurance "github.com/johnazedo/ms-insurance/src/xp/insurance"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
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

// @host 127.0.0.1:8000
// @BasePath /

// @schemes http
func main() {
	simulationController := simulation.GetSimulationController()
	phoneListController := phonelist.GetPhoneListController()
	insuranceController := insurance.GetInsuranceController()

	router := gin.Default()
	router.POST("/simulation", simulationController.GetInsuranceSimulation)
	router.POST("/buy", simulationController.BuyInsurance)
	router.GET("/brands", phoneListController.GetPhoneBrands)
	router.GET("/models", phoneListController.GetPhoneModels)

	router.POST("/cancel", insuranceController.CancelInsurance)
	router.POST("/insurance", insuranceController.GetInsuranceInformation)

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	router.Run("127.0.0.1:8000")
}
