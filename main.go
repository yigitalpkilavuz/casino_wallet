package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	controller "github.com/yigitalpkilavuz/casino_wallet/api/controllers"
	middleware "github.com/yigitalpkilavuz/casino_wallet/api/middlewares"
	config "github.com/yigitalpkilavuz/casino_wallet/config"
	storage "github.com/yigitalpkilavuz/casino_wallet/database"
	logger "github.com/yigitalpkilavuz/casino_wallet/log"
	repository "github.com/yigitalpkilavuz/casino_wallet/repositories"
	service "github.com/yigitalpkilavuz/casino_wallet/services"
)

type App struct {
	Config                config.Config
	Logger                *logrus.Logger
	BaseRepository        repository.BaseRepository
	WalletRepository      repository.WalletRepository
	TransactionRepository repository.TransactionRepository
	BaseService           service.BaseService
	WalletService         service.WalletService
	WalletController      controller.WalletController
}

func NewApp() *App {
	config, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	logger := logger.NewLogger()

	db, err := storage.InitDatabase(config.Database.StorageType, config.Database.ConnectionString)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = storage.RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	baseRepository := repository.NewBaseRepository(db)
	walletRepository := repository.NewWalletRepository(baseRepository)
	transactionRepository := repository.NewTransactionRepository(baseRepository)
	baseService := service.NewBaseService(walletRepository, transactionRepository)
	walletService := service.NewWalletService(baseService)
	walletController := controller.NewWalletController(walletService)

	return &App{
		Config:                config,
		Logger:                logger,
		BaseRepository:        baseRepository,
		WalletRepository:      walletRepository,
		TransactionRepository: transactionRepository,
		BaseService:           baseService,
		WalletService:         walletService,
		WalletController:      walletController,
	}
}

func (app *App) Start() {
	server := gin.Default()
	v1 := server.Group("api/v1")
	app.addRoutes(v1)
	server.Run(":8080")
}

func (app *App) addRoutes(rg *gin.RouterGroup) {
	wallet := rg.Group("/wallet")
	wallet.Use(middleware.ErrorMiddleware())
	wallet.Use(middleware.LoggerMiddleware(app.Logger))
	wallet.Use(middleware.AuthMiddleware())
	wallet.POST("/authenticate", app.WalletController.Authenticate)
	wallet.GET("/:wallet_id/balance", app.WalletController.Balance)
	wallet.POST("/:wallet_id/credit", app.WalletController.Credit)
	wallet.POST("/:wallet_id/debit", app.WalletController.Debit)
}

func main() {
	app := NewApp()
	app.Start()
}
