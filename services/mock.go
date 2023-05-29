package service

import (
	"reflect"
	"time"

	"github.com/stretchr/testify/mock"
	entity "github.com/yigitalpkilavuz/casino_wallet/database/entities"
)

type MockBaseRepository struct {
	mock.Mock
}
type MockWalletRepository struct {
	MockBaseRepository
}

type MockRedisService struct {
	mock.Mock
}

func (m *MockWalletRepository) GetWalletByUsername(username string) (entity.Wallet, error) {
	args := m.Called(username)
	return args.Get(0).(entity.Wallet), nil
}
func (m *MockBaseRepository) Get(id string, out interface{}) error {
	args := m.Called(id, out)
	if args.Get(0) != nil {
		data := reflect.ValueOf(args.Get(0))
		reflect.ValueOf(out).Elem().Set(data)
	}
	return args.Error(1)
}

func (m *MockBaseRepository) Create(data interface{}) error {
	return nil
}

func (m *MockBaseRepository) Update(id string, data interface{}) error {
	return nil
}

func (m *MockBaseRepository) Delete(id string, model interface{}) error {
	return nil
}

func (m MockRedisService) Set(key string, value string, expiration time.Duration) error {
	return nil
}

func (m MockRedisService) Get(key string) (string, error) {
	return "", nil
}

func (m MockRedisService) Clear(key string) error {
	return nil
}

func (m MockRedisService) Exists(key string) (bool, error) {
	return false, nil
}
