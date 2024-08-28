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

type DistrictUseCase struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.DistrictRepository[int, []int]
}

func NewDistrictUseCase(log *zap.Logger, db *gorm.DB, repository *repository.DistrictRepository[int, []int]) *DistrictUseCase {
	return &DistrictUseCase{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *DistrictUseCase) List(ctx context.Context, request model.ListDistrictByIDRequest[[]int]) (*[]entity.District, int64, error) {
	callbackContext := &Context[model.ListDistrictByIDRequest[[]int], []entity.District]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.District](
		callbackContext,
		func(ctx *Context[model.ListDistrictByIDRequest[[]int], []entity.District]) (ContextResult[[]entity.District], error) {
			data := ctx.Data

			where := map[string]interface{}{}
			if data.Request.ID != nil {
				where["id"] = data.Request.ID
			}
			if data.Request.IDCity != nil {
				where["id_city"] = data.Request.IDCity
			}
			if data.Request.IDProvince != nil {
				where["id_province"] = data.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(data.DB, where)
			if err != nil {
				return ContextResult[[]entity.District]{}, err
			}

			return ContextResult[[]entity.District]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *DistrictUseCase) GetById(ctx context.Context, request model.GetDistrictByIDRequest[[]int]) (*[]entity.District, int64, error) {
	callbackContext := &Context[model.GetDistrictByIDRequest[[]int], []entity.District]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.District](
		callbackContext,
		func(ctx *Context[model.GetDistrictByIDRequest[[]int], []entity.District]) (ContextResult[[]entity.District], error) {
			data := ctx.Data

			where := map[string]interface{}{}
			if data.Request.ID != nil {
				where["id"] = data.Request.ID
			}
			if data.Request.IDCity != nil {
				where["id_city"] = data.Request.IDCity
			}
			if data.Request.IDProvince != nil {
				where["id_province"] = data.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(data.DB, where)
			if err != nil {
				return ContextResult[[]entity.District]{}, err
			}

			return ContextResult[[]entity.District]{Collection: collections, Total: total}, nil
		},
	)
}

func (uc *DistrictUseCase) GetFirstById(ctx context.Context, request model.GetDistrictByIDRequest[int]) (**entity.District, int64, error) {
	callbackContext := &Context[model.GetDistrictByIDRequest[int], *entity.District]{Data: *NewContextData(ctx, uc.Log, uc.DB, request)}
	return Wrapper[entity.District](
		callbackContext,
		func(ctx *Context[model.GetDistrictByIDRequest[int], *entity.District]) (ContextResult[*entity.District], error) {
			data := ctx.Data

			id := data.Request.ID
			idCity := data.Request.IDCity
			idProvince := data.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdCityAndIdProvince(data.DB, id, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)
					data.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return ContextResult[*entity.District]{}, apperror.RecordNotFound(errorMessage)
				}

				data.Log.Warn(err.Error())
				return ContextResult[*entity.District]{}, err
			}

			return ContextResult[*entity.District]{Collection: collection, Total: 1}, nil
		},
	)
}
