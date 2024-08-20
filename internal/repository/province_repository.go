package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProvinceRepository struct {
	Repository[entity.Province]
	Log *zap.Logger
}

func NewProvinceRepository(log *zap.Logger) *ProvinceRepository {
	return &ProvinceRepository{
		Log: log,
	}
}

func (r *ProvinceRepository) FindAll(tx *gorm.DB) ([]entity.Province, error) {
	var provinces []entity.Province
	if err := tx.Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
