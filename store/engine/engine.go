package engine

import (
	"context"
	"database/sql"

	"github.com/Dibyajyotid/car-management-backend/models"
)

type EngineStore struct {
	db *sql.DB
}

func new(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) UpdateEngine(stx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {

}
