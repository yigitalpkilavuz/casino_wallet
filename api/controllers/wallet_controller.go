package controller

import (
	"github.com/gin-gonic/gin"
	models "github.com/yigitalpkilavuz/casino_wallet/models"
	service "github.com/yigitalpkilavuz/casino_wallet/services"
)

type WalletController struct {
	walletService service.WalletService
}

func NewWalletController(walletService service.WalletService) WalletController {
	return WalletController{
		walletService: walletService,
	}
}

func (c *WalletController) Authenticate(ctx *gin.Context) {
	req := models.AuthenticateRequest{}
	ctx.BindJSON(&req)
	response, err := c.walletService.Authenticate(req)
	if err.Status > 0 {
		ctx.JSON(err.Status, err)
	} else {
		ctx.JSON(200, response)
	}
}

func (c *WalletController) Balance(ctx *gin.Context) {
	walletId := ctx.Param("wallet_id")
	response, err := c.walletService.Balance(walletId)
	if err.Status > 0 {
		ctx.JSON(err.Status, err)
	} else {
		ctx.JSON(200, response)
	}

}

func (c *WalletController) Credit(ctx *gin.Context) {
	req := models.TransactionRequest{}
	ctx.BindJSON(&req)
	req.WalletId = ctx.Param("wallet_id")
	response, err := c.walletService.Credit(req)
	if err.Status > 0 {
		ctx.JSON(err.Status, err)
	} else {
		ctx.JSON(200, response)
	}
}

func (c *WalletController) Debit(ctx *gin.Context) {
	req := models.TransactionRequest{}
	ctx.BindJSON(&req)
	req.WalletId = ctx.Param("wallet_id")
	response, err := c.walletService.Debit(req)
	if err.Status > 0 {
		ctx.JSON(err.Status, err)
	} else {
		ctx.JSON(200, response)
	}
}
