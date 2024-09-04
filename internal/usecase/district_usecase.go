package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/repository"
	appusecase "github.com/aikuci/go-subdivisions-id/pkg/usecase"
	apperror "github.com/aikuci/go-subdivisions-id/pkg/util/error"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type District struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.District[int, []int]
}

func NewDistrict(log *zap.Logger, db *gorm.DB, repository *repository.District[int, []int]) *District {
	return &District{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *District) List(ctx context.Context, request model.ListDistrictByIDRequest[[]int]) (*[]entity.District, int64, error) {
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.ListDistrictByIDRequest[[]int]]) (*[]entity.District, int64, error) {
			where := map[string]interface{}{}
			if ctx.Request.ID != nil {
				where["id"] = ctx.Request.ID
			}
			if ctx.Request.IDCity != nil {
				where["id_city"] = ctx.Request.IDCity
			}
			if ctx.Request.IDProvince != nil {
				where["id_province"] = ctx.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(ctx.DB, where)
			return &collections, total, err
		},
	)
}

func (uc *District) GetById(ctx context.Context, request model.GetDistrictByIDRequest[[]int]) (*[]entity.District, int64, error) {
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetDistrictByIDRequest[[]int]]) (*[]entity.District, int64, error) {
			where := map[string]interface{}{}
			if ctx.Request.ID != nil {
				where["id"] = ctx.Request.ID
			}
			if ctx.Request.IDCity != nil {
				where["id_city"] = ctx.Request.IDCity
			}
			if ctx.Request.IDProvince != nil {
				where["id_province"] = ctx.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(ctx.DB, where)
			return &collections, total, err
		},
	)
}

func (uc *District) GetFirstById(ctx context.Context, request model.GetDistrictByIDRequest[int]) (*entity.District, int64, error) {
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetDistrictByIDRequest[int]]) (*entity.District, int64, error) {
			id := ctx.Request.ID
			idCity := ctx.Request.IDCity
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdCityAndIdProvince(ctx.DB, id, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID City: %d and ID Province: %d", id, idCity, idProvince)
					ctx.Log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, 0, apperror.RecordNotFound(errorMessage)
				}

				ctx.Log.Warn(err.Error())
				return nil, 0, err
			}

			return &collection, 1, nil
		},
	)
}
