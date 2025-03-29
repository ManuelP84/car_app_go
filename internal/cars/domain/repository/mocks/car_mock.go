package mocks

import (
	"github.com/ManuelP84/car_app_go/internal/cars/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) FindAll() ([]*model.Car, error) {
	args := m.Called()
	return args.Get(0).([]*model.Car), args.Error(1)
}

func (m *MockCarRepository) FindCarById(id string) (*model.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockCarRepository) CreateCar(car *model.Car) (*model.Car, error) {
	args := m.Called(car)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockCarRepository) UpdateCar(car *model.Car) (*model.Car, error) {
	args := m.Called(car)
	return args.Get(0).(*model.Car), args.Error(1)
}
