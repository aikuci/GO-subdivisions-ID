package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
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

func (uc *ProvinceUseCase) List(ctx context.Context, request *model.ListProvinceRequest) ([]model.ProvinceResponse, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	provinces, err := uc.ProvinceRepository.FindAll(tx)
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("error", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ProvinceResponse, len(provinces))
	for i, province := range provinces {
		responses[i] = *mapper.ProvinceToResponse(&province)
	}

	return responses, nil
}

func (uc *ProvinceUseCase) Get(ctx context.Context, request *model.GetProvinceRequest) (*model.ProvinceResponse, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	province := new(entity.Province)
	ID := request.ID
	if err := uc.ProvinceRepository.FindById(tx, province, ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn(err.Error(), zap.String("error", fmt.Sprintf("failed to get province with ID: %d", ID)))
			return nil, fiber.ErrNotFound
		}

		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("error", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return mapper.ProvinceToResponse(province), nil
}
