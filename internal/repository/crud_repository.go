package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CrudRepository struct {
	Repository[entity.Province]
	Log *zap.Logger
}

func NewCrudRepository(log *zap.Logger) *CrudRepository {
	return &CrudRepository{
		Log: log,
	}
}

func (r *CrudRepository) FindAll(tx *gorm.DB) ([]entity.Province, error) {
	var provinces []entity.Province
	if err := tx.Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
