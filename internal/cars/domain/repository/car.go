package repository

import "github.com/ManuelP84/car_app_go/internal/cars/domain/model"

type CarRepository interface {
	FindAll() ([]*model.Car, error)
	FindCarById(id string) (*model.Car, error)
	CreateCar(car *model.Car) (*model.Car, error)
	UpdateCar(car *model.Car) (*model.Car, error)
}
