package car

import (
	"context"
	"database/sql"

	"github.com/Dibyajyotid/car-management-backend/models"
)

type CarStore struct {
	db *sql.DB
}

func new(db *sql.DB) CarStore {
	return CarStore{db: db}
}

func (s CarStore) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car

	query := `SELECT c.id, c.name, c.year, brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range FROM car c LEFT JOIN engine e ON c.enine_id = e.id WHERE c.id=$1`

	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Engine.EngineID,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngineID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCylinders,
		&car.Engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}

		return car, err
	}

	return car, nil
}

func (s CarStore) GetCarByBrand(ctx context.Context, brand string, isEngine bool) {

}

func (s CarStore) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {

}

func (s CarStore) UpdateCar(stx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {

}

func (s CarStore) DeleteCar(ctx context.Context, id string) (models.Car, error) {

}
