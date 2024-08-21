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
	Log        *zap.Logger
	Validate   *validator.Validate
	DB         *gorm.DB
	Repository repository.CruderRepository[entity.Province]
}

func NewProvinceUseCase(logger *zap.Logger, db *gorm.DB, repository repository.CruderRepository[entity.Province],
) *ProvinceUseCase {
	return &ProvinceUseCase{
		Log:        logger,
		DB:         db,
		Repository: repository,
	}
}

func (uc *ProvinceUseCase) List(ctx context.Context, request *model.ListRequest) ([]model.ProvinceResponse, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := uc.Repository.FindAll(tx)
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ProvinceResponse, len(data))
	for i, province := range data {
		responses[i] = *mapper.ProvinceToResponse(&province)
	}

	return responses, nil
}

func (uc *ProvinceUseCase) GetByID(ctx context.Context, request *model.GetByIDRequest) (*model.ProvinceResponse, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	province := new(entity.Province)
	ID := request.ID
	if err := uc.Repository.FindById(tx, province, ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn(err.Error(), zap.String("errorMessage", fmt.Sprintf("failed to get province with ID: %d", ID)))
			return nil, fiber.ErrNotFound
		}

		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	return mapper.ProvinceToResponse(province), nil
}
