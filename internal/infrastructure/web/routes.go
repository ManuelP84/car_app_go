package web

import (
	"database/sql"

	"github.com/ManuelP84/car_app_go/internal/cars/aplication/handler"
	"github.com/ManuelP84/car_app_go/internal/cars/aplication/service"
	"github.com/ManuelP84/car_app_go/internal/infrastructure/database/mysql"
	"github.com/go-chi/chi/v5"
)

func BuildCarsRoutes(rt chi.Router, db *sql.DB) {
	carsRepository := mysql.NewMySQLCarRepository(db)
	carService := service.NewCarService(carsRepository)
	carHandler := handler.NewCarHandler(carService)

	rt.Route("/cars", func(rt chi.Router) {
		rt.Get("/all", carHandler.FindAll)
		rt.Get("/{id}", carHandler.FindCarById)
		rt.Post("/", carHandler.CreateCar)
		rt.Put("/update/{id}", carHandler.UpdateCar)
	})
}
