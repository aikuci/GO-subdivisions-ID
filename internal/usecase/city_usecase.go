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

func (uc *CityUseCase) List(ctx context.Context, request model.ListRequest) ([]entity.City, error) {
	return uc.CrudUseCase.List(ctx, request)
}
func (uc *CityUseCase) GetByID(ctx context.Context, request model.GetByIDRequest[int]) ([]entity.City, error) {
	return uc.CrudUseCase.GetByID(ctx, request)
}
func (uc *CityUseCase) GetByIDs(ctx context.Context, request model.GetByIDRequest[[]int]) ([]entity.City, error) {
	return uc.CrudUseCase.GetByIDs(ctx, request)
}
func (uc *CityUseCase) GetFirstByID(ctx context.Context, request model.GetByIDRequest[int]) (*entity.City, error) {
	return uc.CrudUseCase.GetFirstByID(ctx, request)
}

// Specific UseCase
func (uc *CityUseCase) ListFindByIDAndIDProvince(ctx context.Context, request model.ListCityByIDRequest[[]int]) ([]entity.City, error) {
	useCase := newUseCase[entity.City](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperPlural(
		ctx,
		useCase,
		func(cp *CallbackParam[model.ListCityByIDRequest[[]int]]) ([]entity.City, error) {
			where := map[string]interface{}{}
			if cp.request.ID != nil {
				where["id"] = cp.request.ID
			}
			if cp.request.IDProvince != nil {
				where["id_province"] = cp.request.IDProvince
			}
			return uc.Repository.FindBy(cp.tx, where)
		},
	)
}

func (uc *CityUseCase) GetFindByIDAndIDProvince(ctx context.Context, request model.GetCityByIDRequest[[]int]) ([]entity.City, error) {
	useCase := newUseCase[entity.City](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperPlural(
		ctx,
		useCase,
		func(cp *CallbackParam[model.GetCityByIDRequest[[]int]]) ([]entity.City, error) {
			where := map[string]interface{}{}
			if cp.request.ID != nil {
				where["id"] = cp.request.ID
			}
			if cp.request.IDProvince != nil {
				where["id_province"] = cp.request.IDProvince
			}
			return uc.Repository.FindBy(cp.tx, where)
		},
	)
}

func (uc *CityUseCase) GetFirstByIDAndIDProvince(ctx context.Context, request model.GetCityByIDRequest[int]) (*entity.City, error) {
	useCase := newUseCase[entity.City](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperSingular(
		ctx,
		useCase,
		func(cp *CallbackParam[model.GetCityByIDRequest[int]]) (*entity.City, error) {
			id := cp.request.ID
			idProvince := cp.request.IDProvince
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
		},
	)
}
