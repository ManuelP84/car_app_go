package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ManuelP84/car_app_go/internal/cars/aplication/dto"
	"github.com/ManuelP84/car_app_go/internal/cars/aplication/models"
	"github.com/ManuelP84/car_app_go/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type CarHandler struct {
	carService models.CarService
}

func NewCarHandler(carService models.CarService) *CarHandler {
	return &CarHandler{
		carService: carService,
	}
}

func (h *CarHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	cars, err := h.carService.FindAll()

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if len(cars) == 0 {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"message": "No cars found"})
		return
	}

	response := map[string]interface{}{
		"message": "Cars retrieved successfully",
		"data":    cars,
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *CarHandler) FindCarById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	car, err := h.carService.FindCarById(id)

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return

	}

	if car == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Car not found"})
		return

	}

	response := map[string]interface{}{
		"message": "Car found",
		"data":    car,
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var carDto *dto.CarDto

	err := json.NewDecoder(r.Body).Decode(&carDto)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		return
	}

	err = h.carService.CreateCar(carDto)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Car created"})
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var carDto *dto.CarDto

	err := json.NewDecoder(r.Body).Decode(&carDto)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		return
	}

	if carDto.ID != id {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "ID mismatch"})
		return
	}

	err = h.carService.UpdateCar(carDto)
	if err != nil {
		if err.Error() == "car not found" {
			utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Car not found"})
		} else {
			utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Car updated"})
}
