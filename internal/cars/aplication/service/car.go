package service

import (
	"errors"

	"github.com/ManuelP84/car_app_go/internal/cars/aplication/dto"
	"github.com/ManuelP84/car_app_go/internal/cars/domain/model"
	"github.com/ManuelP84/car_app_go/internal/cars/domain/repository"
)

type CarService struct {
	carRepository repository.CarRepository
}

func NewCarService(carRepository repository.CarRepository) *CarService {
	return &CarService{
		carRepository: carRepository,
	}
}

func (s *CarService) FindAll() ([]*model.Car, error) {
	cars, err := s.carRepository.FindAll()

	//TODO: handle error
	if err != nil {
		return nil, err
	}

	if cars == nil {
		return nil, nil
	}

	//TODO: map the response to a DTO
	return cars, nil
}

func (s *CarService) FindCarById(id string) (*dto.CarDto, error) {
	car, err := s.carRepository.FindCarById(id)

	//TODO: handle error
	if err != nil {
		return nil, err
	}

	//TODO: handle error
	if car == nil {
		return nil, nil
	}

	//TODO: map the response to a DTO
	carDto, _ := dto.MapCarToResponse(*car)
	return &carDto, nil
}

func (s *CarService) CreateCar(carDto *dto.CarDto) error {

	//TODO: validate the car

	car, _ := dto.MapRequestToCar(*carDto)

	_, err := s.carRepository.CreateCar(&car)

	if err != nil {
		return err
	}

	return nil
}

func (s *CarService) UpdateCar(carDto *dto.CarDto) error {
	car, _ := dto.MapRequestToCar(*carDto)

	carDB, err := s.carRepository.FindCarById(car.ID)

	if err != nil || carDB == nil {
		return errors.New("car not found")
	}

	_, err = s.carRepository.UpdateCar(&car)

	if err != nil {
		return errors.New("unable to update the car")
	}

	return nil
}
