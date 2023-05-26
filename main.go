package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/yigitalpkilavuz/casino_wallet/api/controllers"
	conf "github.com/yigitalpkilavuz/casino_wallet/conf"
	service "github.com/yigitalpkilavuz/casino_wallet/services"
)

var (
	walletService      service.WalletService       = service.NewUserService()
	transactionService service.TransactionService  = service.NewTransactionService()
	walletController   controller.WalletController = controller.NewWalletController(walletService, transactionService)
)

func main() {

	config, err := conf.GetConfig()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}

	server := gin.Default()
	// server.Use(middleware.ErrorMiddleware())
	// server.Use(middleware.LoggerMiddleware())
	v1 := server.Group("api/v1")
	AddRoutes(v1)
	server.Run(":8080")
}

func AddRoutes(rg *gin.RouterGroup) {
	wallet := rg.Group("/wallet")
	wallet.GET("/:wallet_id/authenticate", walletController.Authenticate)
	wallet.GET("/:wallet_id/balance", walletController.Balance)
	wallet.POST("/:wallet_id/credit", walletController.Credit)
	wallet.GET("/:wallet_id/debit", walletController.Debit)
}
