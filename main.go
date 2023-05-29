package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	controller "github.com/yigitalpkilavuz/casino_wallet/api/controllers"
	middleware "github.com/yigitalpkilavuz/casino_wallet/api/middlewares"
	"github.com/yigitalpkilavuz/casino_wallet/caching"
	storage "github.com/yigitalpkilavuz/casino_wallet/database"
	config "github.com/yigitalpkilavuz/casino_wallet/framework/config"
	logger "github.com/yigitalpkilavuz/casino_wallet/framework/log"
	repository "github.com/yigitalpkilavuz/casino_wallet/repositories"
	service "github.com/yigitalpkilavuz/casino_wallet/services"
)

// App struct holds all the necessary components of the application.
type App struct {
	Config                config.Config
	Logger                *logrus.Logger
	BaseRepository        repository.BaseRepository
	WalletRepository      repository.WalletRepository
	TransactionRepository repository.TransactionRepository
	BaseService           service.BaseService
	RedisService          caching.RedisService
	WalletService         service.IWalletService
	WalletController      controller.WalletController
}

// NewApp function initializes a new application.
func NewApp() *App {
	// Initialize the configuration
	config, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Initialize logger
	logger := logger.NewLogger()

	// Initialize the database with given configuration
	db, err := storage.InitDatabase(config.Database.StorageType, config.Database.ConnectionString)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run the database migrations
	err = storage.RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories and services
	baseRepository := repository.NewBaseRepository(db)
	walletRepository := repository.NewWalletRepository(baseRepository)
	transactionRepository := repository.NewTransactionRepository(baseRepository)
	baseService := service.NewBaseService(&walletRepository, transactionRepository)
	redisService := caching.NewRedisService(config)
	walletService := service.NewWalletService(baseService, redisService, logger)
	walletController := controller.NewWalletController(&walletService)

	// Return a new app instance
	return &App{
		Config:                config,
		Logger:                logger,
		BaseRepository:        baseRepository,
		WalletRepository:      walletRepository,
		TransactionRepository: transactionRepository,
		BaseService:           baseService,
		RedisService:          redisService,
		WalletService:         &walletService,
		WalletController:      walletController,
	}
}

// Start method starts the Gin server
func (app *App) Start() {
	// Initialize the server
	server := gin.Default()

	// Create API group for version 1
	v1 := server.Group("api/v1")

	// Add routes to the group
	app.addRoutes(v1)

	// Run the server on port 8080
	server.Run(":8080")
}

// addRoutes method adds routes to a given router group
func (app *App) addRoutes(rg *gin.RouterGroup) {
	// Create wallet group under given router group
	wallet := rg.Group("/wallet")

	// Use middlewares for the wallet group
	wallet.Use(middleware.ErrorMiddleware(app.Logger))
	wallet.Use(middleware.LoggerMiddleware(app.Logger))
	wallet.Use(middleware.AuthMiddleware())

	// Add routes for the wallet group
	wallet.POST("/authenticate", app.WalletController.Authenticate)
	wallet.GET("/:wallet_id/balance", app.WalletController.Balance)
	wallet.POST("/:wallet_id/credit", app.WalletController.Credit)
	wallet.POST("/:wallet_id/debit", app.WalletController.Debit)
}

func main() {
	app := NewApp()
	app.Start()
}
