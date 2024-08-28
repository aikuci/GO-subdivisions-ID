package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

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

func (uc *CityUseCase) List(ctx context.Context, request model.ListCityByIDRequest[[]int]) (*[]entity.City, int64, error) {
	callbackContext := &Context[model.ListCityByIDRequest[[]int], []entity.City]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.City](
		callbackContext,
		func(ctx *Context[model.ListCityByIDRequest[[]int], []entity.City]) (ContextResult[[]entity.City], error) {
			data := ctx.Data

			where := map[string]interface{}{}
			if data.Request.ID != nil {
				where["id"] = data.Request.ID
			}
			if data.Request.IDProvince != nil {
				where["id_province"] = data.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(data.DB, where)
			if err != nil {
				return ContextResult[[]entity.City]{}, err
			}

			return ContextResult[[]entity.City]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *CityUseCase) GetById(ctx context.Context, request model.GetCityByIDRequest[[]int]) (*[]entity.City, int64, error) {
	callbackContext := &Context[model.GetCityByIDRequest[[]int], []entity.City]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.City](
		callbackContext,
		func(ctx *Context[model.GetCityByIDRequest[[]int], []entity.City]) (ContextResult[[]entity.City], error) {
			data := ctx.Data

			where := map[string]interface{}{}
			if data.Request.ID != nil {
				where["id"] = data.Request.ID
			}
			if data.Request.IDProvince != nil {
				where["id_province"] = data.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(data.DB, where)
			if err != nil {
				return ContextResult[[]entity.City]{}, err
			}

			return ContextResult[[]entity.City]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *CityUseCase) GetFirstById(ctx context.Context, request model.GetCityByIDRequest[int]) (**entity.City, int64, error) {
	callbackContext := &Context[model.GetCityByIDRequest[int], *entity.City]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.City](
		callbackContext,
		func(ctx *Context[model.GetCityByIDRequest[int], *entity.City]) (ContextResult[*entity.City], error) {
			data := ctx.Data

			id := data.Request.ID
			idProvince := data.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdProvince(data.DB, id, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)
					data.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return ContextResult[*entity.City]{}, apperror.RecordNotFound(errorMessage)
				}

				data.Log.Warn(err.Error())
				return ContextResult[*entity.City]{}, err
			}

			return ContextResult[*entity.City]{Collection: collection, Total: 1}, nil
		},
	)
}
