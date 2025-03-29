package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ManuelP84/car_app_go/internal/infrastructure/app"
	"github.com/ManuelP84/car_app_go/internal/infrastructure/database/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServerChi(config *app.Config) *ServerChi {
	return &ServerChi{
		cfg: config,
	}
}

type ServerChi struct {
	cfg *app.Config
}

func (s *ServerChi) Run() (err error) {
	// setup db
	db := mysql.MySQLConnection(s.cfg)
	defer db.Close()

	// router
	rt := chi.NewRouter()

	// middelwares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// endpoints
	rt.Route("/api/v1", func(r chi.Router) {

		// domains endpoints
		BuildCarsRoutes(r, db)

	})

	// create server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", s.cfg.Port),
		Handler:           rt,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// run server
	err = srv.ListenAndServe()
	if err != nil {
		return
	}

	return
}
