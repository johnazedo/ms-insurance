package main

import (
	"github.com/gin-gonic/gin"
	simulation "github.com/johnazedo/ms-insurance/src/offer/simulation"
	phonelist "github.com/johnazedo/ms-insurance/src/offer/phonelist"
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

// @host 172.17.136.193:8000
// @BasePath /

// @schemes http
func main() {
	simulationController := simulation.Controller{}
	phonelistController := phonelist.Controller{}

	router := gin.Default()
	router.POST("/simulation", simulationController.GetInsuranceSimulation)
	router.POST("/buy", simulationController.BuyInsurance)
	
	router.GET("/brands", phonelistController.GetPhoneBrands)
	router.GET("/models", phonelistController.GetPhoneModels)

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	
	router.Run("172.17.136.193:8000")
}