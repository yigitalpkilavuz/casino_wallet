package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	service "github.com/yigitalpkilavuz/casino_wallet/services"
)

type WalletController interface {
	Authenticate(ctx *gin.Context)
	Balance(ctx *gin.Context)
	Credit(ctx *gin.Context)
	Debit(ctx *gin.Context)
}

type walletController struct {
	walletService      service.WalletService
	transactionService service.TransactionService
}

func NewWalletController(walletService service.WalletService, transactionService service.TransactionService) WalletController {
	return &walletController{
		walletService:      walletService,
		transactionService: transactionService,
	}
}

func (c *walletController) Authenticate(ctx *gin.Context) {
	fmt.Print("test")
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func (c *walletController) Balance(ctx *gin.Context) {
	response, err := c.walletService.Find(1)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{
		"balance": response.Balance,
	})

}

func (c *walletController) Credit(ctx *gin.Context) {
}

func (c *walletController) Debit(ctx *gin.Context) {
}
