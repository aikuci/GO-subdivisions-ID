package usecase

import (
	"context"

	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/middleware/requestid"
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CityUseCase struct {
	CrudUseCase CrudUseCase[entity.City, model.CityResponse] // embedded

	Log        *zap.Logger
	Validate   *validator.Validate
	DB         *gorm.DB
	Repository repository.CrudRepository[entity.City]
	Mapper     mapper.CruderMapper[entity.City, model.CityResponse]
}

func NewCityUseCase(log *zap.Logger, db *gorm.DB, repository *repository.CrudRepository[entity.City], mapper mapper.CruderMapper[entity.City, model.CityResponse]) *CityUseCase {
	crudUseCase := NewCrudUseCase(log, db, repository, mapper)

	return &CityUseCase{
		CrudUseCase: *crudUseCase,

		Log:        log,
		DB:         db,
		Repository: *repository,
		Mapper:     mapper,
	}
}

func (uc *CityUseCase) List(ctx context.Context, request *model.ListRequest) ([]model.CityResponse, error) {
	return uc.CrudUseCase.List(ctx, request)
}

func (uc *CityUseCase) GetByID(ctx context.Context, request *model.GetByIDRequest) (*model.CityResponse, error) {
	return uc.CrudUseCase.GetByID(ctx, request)
}

func (uc *CityUseCase) GetByIDReturnCollections(ctx context.Context, request *model.GetByIDRequest) ([]model.CityResponse, error) {
	logger := uc.Log.With(zap.String(string("requestid"), requestid.FromContext(ctx)))

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	collections, err := uc.Repository.FindBy(tx, map[string]interface{}{"id": request.ID})
	if err != nil {
		logger.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		logger.Warn(err.Error(), zap.String("errorMessage", "failed to commit transaction"))
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.CityResponse, len(collections))
	for i, collection := range collections {
		responses[i] = *uc.Mapper.ModelToResponse(&collection)
	}

	return responses, nil
}
