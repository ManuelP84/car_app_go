package main

import (
	"fmt"

	"github.com/ManuelP84/car_app_go/internal/infrastructure/app"
	"github.com/ManuelP84/car_app_go/internal/infrastructure/web"

)

// @title Car API
// @version 1.0
// @description This is a sample server for a car API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := app.LoadConfig()

	server := web.NewServerChi(cfg)

	fmt.Println("Server running on port: ", cfg.Port)
	if err := server.Run(); err != nil {
		fmt.Println()
		return
	}

}
