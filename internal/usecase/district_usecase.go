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
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.ListDistrictByIDRequest[[]int]]) (*[]entity.District, int64, error) {
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

func (uc *DistrictUseCase) GetById(ctx context.Context, request model.GetDistrictByIDRequest[[]int]) (*[]entity.District, int64, error) {
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetDistrictByIDRequest[[]int]]) (*[]entity.District, int64, error) {
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

func (uc *DistrictUseCase) GetFirstById(ctx context.Context, request model.GetDistrictByIDRequest[int]) (*entity.District, int64, error) {
	return appusecase.Wrapper[entity.District](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetDistrictByIDRequest[int]]) (*entity.District, int64, error) {
			id := ctx.Request.ID
			idCity := ctx.Request.IDCity
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdCityAndIdProvince(ctx.DB, id, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince)
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
