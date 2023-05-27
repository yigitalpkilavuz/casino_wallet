package service

// import (
// 	"os"
// 	"testing"

// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	entity "github.com/yigitalpkilavuz/casino_wallet/database/entity"
// 	model "github.com/yigitalpkilavuz/casino_wallet/model"
// 	service "github.com/yigitalpkilavuz/casino_wallet/service"
// )

// func TestWalletService_Authenticate(t *testing.T) {
// 	// Create a test instance of the WalletService
// 	walletRepository := &mockWalletRepository{}
// 	baseService := service.BaseService{WalletRepository: walletRepository}
// 	walletService := service.NewWalletService(baseService)

// 	// Define the test input
// 	request := model.AuthenticateRequest{
// 		Username: "testuser",
// 		Password: "testpassword",
// 	}

// 	// Define the expected output
// 	expectedResponse := model.AuthenticateResponse{
// 		Username: "testuser",
// 		Balance:  decimal.Zero,
// 		Token:    "testtoken",
// 	}
// 	expectedError := model.ErrorResponse{}

// 	// Set up the mock WalletRepository
// 	wallet := entity.Wallet{
// 		Username: "testuser",
// 		Password: "testpassword",
// 		Balance:  decimal.Zero,
// 	}
// 	walletRepository.On("GetWalletByUsername", request.Username).Return(wallet, nil)
// 	auth.On("CreateToken", request.Username).Return("testtoken", nil)

// 	// Call the Authenticate method
// 	response, err := walletService.Authenticate(request)

// 	// Assert the results
// 	assert.Equal(t, expectedResponse, response)
// 	assert.Equal(t, expectedError, err)
// }

// func TestWalletService_Balance(t *testing.T) {
// 	// Create a test instance of the WalletService
// 	walletRepository := &mockWalletRepository{}
// 	baseService := service.BaseService{WalletRepository: walletRepository}
// 	walletService := service.NewWalletService(baseService)

// 	// Define the test input
// 	walletID := "testwalletid"

// 	// Define the expected output
// 	expectedResponse := model.BalanceResponse{
// 		Username: "testuser",
// 		Balance:  decimal.Zero,
// 	}
// 	expectedError := model.ErrorResponse{}

// 	// Set up the mock WalletRepository
// 	wallet := entity.Wallet{
// 		Username: "testuser",
// 		Balance:  decimal.Zero,
// 	}
// 	walletRepository.On("Get", walletID, &wallet).Return(nil)

// 	// Call the Balance method
// 	response, err := walletService.Balance(walletID)

// 	// Assert the results
// 	assert.Equal(t, expectedResponse, response)
// 	assert.Equal(t, expectedError, err)
// }

// // Define a mock implementation of the WalletRepository interface
// type mockWalletRepository struct {
// 	mock.Mock
// }

// func (m *mockWalletRepository) GetWalletByUsername(username string) (entity.Wallet, error) {
// 	args := m.Called(username)
// 	return args.Get(0).(entity.Wallet), args.Error(1)
// }

// func (m *mockWalletRepository) Get(id string, out interface{}) error {
// 	args := m.Called(id, out)
// 	return args.Error(0)
// }

// // ... Implement the remaining methods of the WalletRepository interface in a similar way ...

// // Define a mock implementation of the Auth package
// type mockAuth struct {
// 	mock.Mock
// }

// func (m *mockAuth) CreateToken(username string) (string, error) {
// 	args := m.Called(username)
// 	return args.String(0), args.Error(1)
// }

// // ... Implement the remaining methods of the Auth package in a similar way ...
// func TestMain(m *testing.M) {
// 	// Run the tests
// 	os.Exit(m.Run())
// }
