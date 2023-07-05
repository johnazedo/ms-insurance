package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/johnazedo/ms-insurance/docs"
	di "github.com/johnazedo/ms-insurance/src/di"
	infra "github.com/johnazedo/ms-insurance/src/infra"
	"github.com/johnazedo/ms-insurance/src/logs"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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

// @host 0.0.0.0:8000
// @BasePath /

// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Config rabbitMQ
	mquser := os.Getenv("RABBITMQ_DEFAULT_USER")
	mqpassword := os.Getenv("RABBITMQ_DEFAULT_PASS")
	mqhost := os.Getenv("RABBITMQ_SERVER_URL")
	mqport := os.Getenv("RABBITMQ_PORT")

	mqconnection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", mquser, mqpassword, mqhost, mqport))
	if err != nil {
		// panic(err)
	}
	defer mqconnection.Close()

	logger := logs.LogServiceImpl{
		Config: &logs.Config{
			Microservice: os.Getenv("SERVER_NAME"),
			Thread:       "Main",
			Context:      "Default",
			IP:           host,
			Key:          "vascobank.logs",
		},
		RabbitMQ: mqconnection,
	}

	// Config database
	db := infra.OpenDB()
	infra.AutoMigrate(db)

	// Confif router

	// recurrence := di.GetRecurrenceUseCase(db, &logger)
	// go recurrence.Invoke()

	router := gin.Default()
	ServeRoutes(router, &logger, db)

	router.Run(fmt.Sprintf("%s:%s", host, port))
}

func ServeRoutes(router *gin.Engine, logger logs.LogService, db *gorm.DB) {
	simulationController := di.GetSimulationController(db)
	phoneListController := di.GetPhoneListController(db)
	insuranceController := di.GetInsuranceController(db, logger)

	router.POST("/simulation", simulationController.GetInsuranceSimulation)
	router.POST("/buy", simulationController.BuyInsurance)
	router.GET("/brands", phoneListController.GetPhoneBrands)
	router.GET("/models", phoneListController.GetPhoneModels)

	router.POST("/cancel", insuranceController.CancelInsurance)
	router.POST("/insurance", insuranceController.GetInsuranceInformation)

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
}
