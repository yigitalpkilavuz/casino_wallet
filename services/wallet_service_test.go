package service

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"
	logger "github.com/yigitalpkilavuz/casino_wallet/framework/log"
	model "github.com/yigitalpkilavuz/casino_wallet/models"
)

func TestAuthenticate(t *testing.T) {
	mockWalletRepo := new(MockWalletRepository)
	mockRedisService := new(MockRedisService)

	testService := WalletService{
		BaseService: BaseService{
			walletRepository: mockWalletRepo,
		},
		Cache:  mockRedisService,
		Logger: logger.NewLogger(),
	}
	t.Run("success", func(t *testing.T) {
		testReq := model.AuthenticateRequest{Username: "test", Password: "password"}
		expectedWallet := entity.Wallet{Username: "test", Password: "password", Balance: decimal.NewFromInt(100)}
		mockWalletRepo.On("GetWalletByUsername", "test").Return(expectedWallet, nil)

		response, err := testService.Authenticate(testReq)
		assert.Equal(t, err, model.ErrorResponse{})
		assert.Equal(t, "test", response.Username)
		assert.Equal(t, decimal.NewFromInt(100), response.Balance)

		mockWalletRepo.AssertExpectations(t)
	})

}

func TestBalance(t *testing.T) {
	mockWalletRepo := new(MockWalletRepository)

	mockRedisService := new(MockRedisService)

	testService := WalletService{
		BaseService: BaseService{
			walletRepository: mockWalletRepo,
		},
		Cache:  mockRedisService,
		Logger: logger.NewLogger(),
	}

	t.Run("success", func(t *testing.T) {
		mockWallet := entity.Wallet{Username: "TestUser", Balance: decimal.NewFromInt(100)}

		mockWalletRepo.On("Get", mock.AnythingOfType("string"), mock.AnythingOfType("*entity.Wallet")).Return(entity.Wallet{Username: "TestUser", Balance: decimal.NewFromInt(100)}, nil)

		resp, err := testService.Balance("1")

		assert.Equal(t, mockWallet.Username, resp.Username)
		assert.Equal(t, mockWallet.Balance, resp.Balance)
		assert.Equal(t, model.ErrorResponse{}, err)

		mockWalletRepo.AssertExpectations(t)
	})
}

// func TestCredit(t *testing.T) {
// 	mockWalletRepo := new(MockWalletRepository)
// 	mockCache := new(MockRedisService)
// 	logger := logger.NewLogger()

// 	testService := WalletService{
// 		BaseService: BaseService{
// 			walletRepository: mockWalletRepo,
// 		},
// 		Cache:  mockCache,
// 		Logger: logger,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		req := model.TransactionRequest{
// 			WalletId: "1",
// 			Amount:   decimal.NewFromInt(100),
// 		}

// 		wallet := entity.Wallet{
// 			Balance: decimal.NewFromInt(500),
// 		}

// 		mockCache.On("Get", req.WalletId).Return("", nil)
// 		mockWalletRepo.On("Get", req.WalletId, mock.AnythingOfType("*entity.Wallet")).Return(nil, nil).Run(func(args mock.Arguments) {
// 			walletArg := args.Get(1).(*entity.Wallet)
// 			*walletArg = wallet
// 		})

// 		mockWalletRepo.On("Create", mock.AnythingOfType("*entity.Transaction")).Return(nil)
// 		mockCache.On("Clear", req.WalletId).Return(nil)
// 		mockCache.On("Set", req.WalletId, mock.AnythingOfType("string"), mock.Anything).Return(nil)

// 		response, err := testService.Credit(req)

// 		assert.Equal(t, model.TransactionResponse{
// 			Username:      wallet.Username,
// 			Balance:       wallet.Balance.Add(req.Amount),
// 			TransactionID: 0,
// 		}, response)

// 		assert.Equal(t, model.ErrorResponse{}, err)

// 		mockCache.AssertExpectations(t)
// 		mockWalletRepo.AssertExpectations(t)
// 	})
// }

// func TestDebit(t *testing.T) {
// 	mockWalletRepo := new(MockWalletRepository)
// 	mockCache := new(MockRedisService)
// 	logger := logger.NewLogger()

// 	testService := WalletService{
// 		BaseService: BaseService{
// 			walletRepository: mockWalletRepo,
// 		},
// 		Cache:  mockCache,
// 		Logger: logger,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		req := model.TransactionRequest{
// 			WalletId: "1",
// 			Amount:   decimal.NewFromInt(100),
// 		}

// 		wallet := entity.Wallet{
// 			Balance: decimal.NewFromInt(500),
// 		}

// 		mockCache.On("Get", req.WalletId).Return("", nil)
// 		mockWalletRepo.On("Get", req.WalletId, mock.AnythingOfType("*entity.Wallet")).Return(nil).Run(func(args mock.Arguments) {
// 			walletArg := args.Get(1).(*entity.Wallet)
// 			*walletArg = wallet
// 		})

// 		mockWalletRepo.On("Create", mock.AnythingOfType("*entity.Transaction")).Return(nil)
// 		mockCache.On("Clear", req.WalletId).Return(nil)
// 		mockCache.On("Set", req.WalletId, mock.AnythingOfType("string"), mock.Anything).Return(nil)

// 		response, err := testService.Debit(req)

// 		assert.Equal(t, model.TransactionResponse{
// 			Username:      wallet.Username,
// 			Balance:       wallet.Balance.Sub(req.Amount),
// 			TransactionID: 0,
// 		}, response)

// 		assert.Equal(t, model.ErrorResponse{}, err)

// 		mockCache.AssertExpectations(t)
// 		mockWalletRepo.AssertExpectations(t)
// 	})

// }
