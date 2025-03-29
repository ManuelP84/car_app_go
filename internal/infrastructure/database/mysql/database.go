package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ManuelP84/car_app_go/internal/infrastructure/app"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection(config *app.Config) *sql.DB {
	stringConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("Database connected")
	return db
}
