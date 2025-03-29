package models

import "net/http"

type CarHandler interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindCarById(w http.ResponseWriter, r *http.Request)
	CreateCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
}
