package usecase

import (
	"context"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProvinceUseCase struct {
	Log                *zap.Logger
	Validate           *validator.Validate
	DB                 *gorm.DB
	ProvinceRepository *repository.ProvinceRepository
}

func NewProvinceUseCase(logger *zap.Logger, db *gorm.DB, provinceRepository *repository.ProvinceRepository,
) *ProvinceUseCase {
	return &ProvinceUseCase{
		Log:                logger,
		DB:                 db,
		ProvinceRepository: provinceRepository,
	}
}

func (uc *ProvinceUseCase) Get(ctx context.Context, request *model.GetProvinceRequest) (*model.ProvinceResponse, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	province := new(entity.Province)
	ID := request.ID
	if err := uc.ProvinceRepository.FindById(tx, province, ID); err != nil {
		message := fmt.Sprintf("failed to get province with ID: %d", ID)
		uc.Log.Warn(err.Error(), zap.String("error", message))
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warn(err.Error(), zap.String("error", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return mapper.ProvinceToResponse(province), nil
}
