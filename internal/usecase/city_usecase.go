package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	apperror "github.com/aikuci/go-subdivisions-id/internal/pkg/error"
	"github.com/aikuci/go-subdivisions-id/internal/repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CityUseCase struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.CityRepository[int, []int]
}

func NewCityUseCase(log *zap.Logger, db *gorm.DB, repository *repository.CityRepository[int, []int]) *CityUseCase {
	return &CityUseCase{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *CityUseCase) List(ctx context.Context, request model.ListCityByIDRequest[[]int]) ([]entity.City, error) {
	useCase := newUseCase[entity.City](uc.Log, uc.DB, request)

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
func (uc *CityUseCase) GetById(ctx context.Context, request model.GetCityByIDRequest[[]int]) ([]entity.City, error) {
	useCase := newUseCase[entity.City](uc.Log, uc.DB, request)

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
func (uc *CityUseCase) GetFirstById(ctx context.Context, request model.GetCityByIDRequest[int]) (*entity.City, error) {
	useCase := newUseCase[entity.City](uc.Log, uc.DB, request)

	return wrapperSingular(
		ctx,
		useCase,
		func(cp *CallbackParam[model.GetCityByIDRequest[int]]) (*entity.City, error) {
			id := cp.request.ID
			idProvince := cp.request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdProvince(cp.tx, id, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)
					cp.log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, apperror.RecordNotFound(errorMessage)
				}

				cp.log.Warn(err.Error())
				return nil, err
			}

			return collection, nil
		},
	)
}
