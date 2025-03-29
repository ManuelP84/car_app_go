package mysql

import (
	"database/sql"
	"log"

	"github.com/ManuelP84/car_app_go/internal/cars/domain/model"
)

type MySQLCarRepository struct {
	db *sql.DB
}

func NewMySQLCarRepository(db *sql.DB) *MySQLCarRepository {
	return &MySQLCarRepository{
		db: db,
	}
}

func (r *MySQLCarRepository) FindAll() ([]*model.Car, error) {
	query := `SELECT * FROM cars`
	rows, err := r.db.Query(query)

	// TODO: handle error
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var cars []*model.Car
	for rows.Next() {
		car := &model.Car{}
		err := rows.Scan(&car.ID, &car.Model, &car.Make, &car.Package, &car.Color, &car.FabricationYear, &car.Category, &car.Mileage, &car.Price)

		// TODO: handle error
		if err != nil {
			log.Println("Error scanning row: %w", err)
			return nil, err
		}
		cars = append(cars, car)
	}

	// TODO: handle error
	if err = rows.Err(); err != nil {
		log.Println("Error scanning row: %w", err)
		return nil, err
	}

	return cars, nil
}

func (r *MySQLCarRepository) FindCarById(id string) (*model.Car, error) {
	query := `SELECT * FROM cars WHERE id = ?`
	row := r.db.QueryRow(query, id)

	// TODO: handle error
	if err := row.Err(); err != nil {
		log.Printf("Error querying row:%s", err.Error())
		return nil, err
	}

	car := &model.Car{}
	err := row.Scan(&car.ID, &car.Model, &car.Make, &car.Package, &car.Color, &car.FabricationYear, &car.Category, &car.Mileage, &car.Price)

	// TODO: handle error
	if err != nil {
		log.Printf("error:%s", err.Error())
		return nil, err
	}

	return car, nil
}

func (r *MySQLCarRepository) CreateCar(car *model.Car) (*model.Car, error) {
	query := `INSERT INTO cars (id, model, make, package, color, fabricationYear, category, mileage, price) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, car.ID, car.Model, car.Make, car.Package, car.Color, car.FabricationYear, car.Category, car.Mileage, car.Price)

	// TODO: handle error
	if err != nil {
		log.Printf("Error in excec: %s", err.Error())
		return nil, err
	}

	_, err = result.LastInsertId()

	// TODO: handle error
	if err != nil {
		log.Printf("error:%s", err.Error())
		return nil, err
	}
	return car, nil
}

func (r *MySQLCarRepository) UpdateCar(car *model.Car) (*model.Car, error) {
	query := `UPDATE cars SET model = ?, make = ?, package = ?, color = ?, fabricationYear = ?, category = ?, mileage = ?, price = ? WHERE id = ?`
	result, err := r.db.Exec(query, car.Model, car.Make, car.Package, car.Color, car.FabricationYear, car.Category, car.Mileage, car.Price, car.ID)

	// TODO: handle error
	if err != nil {
		log.Printf("Error in excec: %s", err.Error())
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()

	// TODO: handle error
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return car, nil
}
