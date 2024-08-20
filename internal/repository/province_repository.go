package repository

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"go.uber.org/zap"
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
