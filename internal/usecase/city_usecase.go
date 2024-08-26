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
	return wrapperPlural(
		newUseCase[entity.City](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.ListCityByIDRequest[[]int]]) ([]entity.City, error) {
			where := map[string]interface{}{}
			if ca.request.ID != nil {
				where["id"] = ca.request.ID
			}
			if ca.request.IDProvince != nil {
				where["id_province"] = ca.request.IDProvince
			}
			return uc.Repository.FindBy(ca.tx, where)
		},
	)
}
func (uc *CityUseCase) GetById(ctx context.Context, request model.GetCityByIDRequest[[]int]) ([]entity.City, error) {
	return wrapperPlural(
		newUseCase[entity.City](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetCityByIDRequest[[]int]]) ([]entity.City, error) {
			where := map[string]interface{}{}
			if ca.request.ID != nil {
				where["id"] = ca.request.ID
			}
			if ca.request.IDProvince != nil {
				where["id_province"] = ca.request.IDProvince
			}
			return uc.Repository.FindBy(ca.tx, where)
		},
	)
}
func (uc *CityUseCase) GetFirstById(ctx context.Context, request model.GetCityByIDRequest[int]) (*entity.City, error) {
	return wrapperSingular(
		newUseCase[entity.City](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetCityByIDRequest[int]]) (*entity.City, error) {
			id := ca.request.ID
			idProvince := ca.request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdProvince(ca.tx, id, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)
					ca.log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, apperror.RecordNotFound(errorMessage)
				}

				ca.log.Warn(err.Error())
				return nil, err
			}

			return collection, nil
		},
	)
}
