package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CityUseCase struct {
	CrudUseCase CrudUseCase[entity.City] // embedded

	Repository repository.CityRepository[int, []int]
}

func NewCityUseCase(logger *zap.Logger, db *gorm.DB, repository *repository.CityRepository[int, []int], mapper mapper.CruderMapper[entity.City, model.CityResponse]) *CityUseCase {
	crudUseCase := NewCrudUseCase(logger, db, repository)

	return &CityUseCase{
		CrudUseCase: *crudUseCase,

		Repository: *repository,
	}
}

func (uc *CityUseCase) List(ctx context.Context, request *model.ListRequest) ([]entity.City, error) {
	return uc.CrudUseCase.List(ctx, request)
}

func (uc *CityUseCase) GetByID(ctx context.Context, request *model.GetByIDRequest[int]) ([]entity.City, error) {
	return uc.CrudUseCase.GetByID(ctx, request)
}

func (uc *CityUseCase) GetByIDs(ctx context.Context, request *model.GetByIDRequest[[]int]) ([]entity.City, error) {
	return uc.CrudUseCase.GetByIDs(ctx, request)
}

func (uc *CityUseCase) GetFirstByID(ctx context.Context, request *model.GetByIDRequest[int]) (*entity.City, error) {
	return uc.CrudUseCase.GetFirstByID(ctx, request)
}

// Specific UseCase
func (uc *CityUseCase) GetByIDProvince(ctx context.Context, request *model.GetCityByIDProvinceRequest[int]) ([]entity.City, error) {
	useCase := NewUseCase[entity.City](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return WrapperPlural(
		ctx,
		useCase,
		uc.getByIdProvinceFn,
	)
}
func (uc *CityUseCase) getByIdProvinceFn(cp *CallbackParam[*model.GetCityByIDProvinceRequest[int]]) ([]entity.City, error) {
	idProvince := cp.request.IDProvince
	collections, err := uc.Repository.FindByIdProvince(cp.tx, idProvince)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cp.log.Warn(err.Error(), zap.String("errorMessage", fmt.Sprintf("failed to get cities data with ID Province: %d", idProvince)))
			return nil, fiber.ErrNotFound
		}

		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collections, nil
}

func (uc *CityUseCase) GetFirstByIDAndIDProvince(ctx context.Context, request *model.GetCityByIDRequest[int]) (*entity.City, error) {
	useCase := NewUseCase[entity.City](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return WrapperSingular(
		ctx,
		useCase,
		uc.getFirstByIDAndIDProvinceFn,
	)
}
func (uc *CityUseCase) getFirstByIDAndIDProvinceFn(cp *CallbackParam[*model.GetCityByIDRequest[int]]) (*entity.City, error) {
	idProvince := cp.request.IDProvince
	id := cp.request.ID
	collection, err := uc.Repository.FirstByIdAndIdProvince(cp.tx, id, idProvince)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cp.log.Warn(err.Error(), zap.String("errorMessage", fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)))
			return nil, fiber.ErrNotFound
		}

		cp.log.Warn(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return collection, nil
}
