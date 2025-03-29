package dto

import "github.com/ManuelP84/car_app_go/internal/cars/domain/model"

type CarDto struct {
	ID              string
	Model           string
	Make            string
	Package         string
	Color           string
	FabricationYear int
	Category        string
	Mileage         int
	Price           int
}

func MapRequestToCar(request CarDto) (model.Car, error) {
	//TODO: ADD validation and return error
	return model.Car{
		ID:              request.ID,
		Model:           request.Model,
		Make:            request.Make,
		Package:         request.Package,
		Color:           request.Color,
		FabricationYear: request.FabricationYear,
		Category:        request.Category,
		Mileage:         request.Mileage,
		Price:           request.Price,
	}, nil
}

func MapCarToResponse(car model.Car) (CarDto, error) {
	//TODO: ADD validation and return error
	return CarDto{
		ID:              car.ID,
		Model:           car.Model,
		Make:            car.Make,
		Package:         car.Package,
		Color:           car.Color,
		FabricationYear: car.FabricationYear,
		Category:        car.Category,
		Mileage:         car.Mileage,
		Price:           car.Price,
	}, nil
}
