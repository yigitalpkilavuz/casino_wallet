package controller

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	models "github.com/yigitalpkilavuz/casino_wallet/models"
// 	service "github.com/yigitalpkilavuz/casino_wallet/services/mocks"
// )

// func TestWalletController_Authenticate(t *testing.T) {
// 	// Set Gin to Test Mode
// 	gin.SetMode(gin.TestMode)

// 	// Setup your controller, mock services and router
// 	mockWalletService := &service.MockWalletService{}
// 	controller := NewWalletController(mockWalletService)
// 	router := gin.Default()
// 	router.POST("/authenticate", controller.Authenticate)

// 	// Mock service functions and set expectations
// 	mockWalletService.("Authenticate", mock.Anything).Return(models.AuthenticateResponse{Username: "test", Balance: decimal.NewFromInt(100)}, models.ErrorResponse{})

// 	// Create a request to send to the route
// 	reqBody := strings.NewReader(`{"username": "test", "password": "password"}`)
// 	req, err := http.NewRequest(http.MethodPost, "/authenticate", reqBody)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v\n", err)
// 	}

// 	// Record HTTP responses
// 	w := httptest.NewRecorder()

// 	// Create HTTP test context and pass it to the router (controller)
// 	router.ServeHTTP(w, req)

// 	// Check the HTTP Status Code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Compare the response body
// 	assert.Contains(t, w.Body.String(), `"username":"test"`)
// 	assert.Contains(t, w.Body.String(), `"balance":1000`)
// }
