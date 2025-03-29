package service_test

import (
	"testing"

	"github.com/ManuelP84/car_app_go/internal/cars/aplication/dto"
	"github.com/ManuelP84/car_app_go/internal/cars/aplication/service"
	"github.com/ManuelP84/car_app_go/internal/cars/domain/model"
	"github.com/ManuelP84/car_app_go/internal/cars/domain/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	service := service.NewCarService(mockRepo)

	cars := []*model.Car{
		{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000},
		{ID: "2", Model: "Model S", Make: "Tesla", Price: 90000},
	}

	mockRepo.On("FindAll").Return(cars, nil)

	result, err := service.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0].Model, "Model X")

	mockRepo.AssertExpectations(t)
}

func TestFindCarById_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	service := service.NewCarService(mockRepo)

	car := &model.Car{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000}
	mockRepo.On("FindCarById", "1").Return(car, nil)

	//ACT
	result, err := service.FindCarById("1")

	//ASSERT
	assert.Nil(t, err)
	assert.Equal(t, result.ID, "1")
	assert.Equal(t, result.Model, "Model X")

	mockRepo.AssertExpectations(t)

}

func TestCreateCar_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	service := service.NewCarService(mockRepo)

	carDto := &dto.CarDto{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000}
	car := &model.Car{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000}

	mockRepo.On("CreateCar", car).Return(car, nil)

	err := service.CreateCar(carDto)

	// Verificar que no hubo error
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCar_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	service := service.NewCarService(mockRepo)

	carDto := &dto.CarDto{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000}
	car := &model.Car{ID: "1", Model: "Model X", Make: "Tesla", Price: 100000}

	mockRepo.On("FindCarById", "1").Return(car, nil)
	mockRepo.On("UpdateCar", car).Return(car, nil)

	err := service.UpdateCar(carDto)

	// Verificar que no hubo error
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
