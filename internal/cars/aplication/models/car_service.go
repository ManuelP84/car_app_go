package models

import (
	"github.com/ManuelP84/car_app_go/internal/cars/aplication/dto"
	domain "github.com/ManuelP84/car_app_go/internal/cars/domain/model"
)

type CarService interface {
	FindAll() ([]*domain.Car, error)
	FindCarById(id string) (*dto.CarDto, error)
	CreateCar(car *dto.CarDto) error
	UpdateCar(car *dto.CarDto) error
}
