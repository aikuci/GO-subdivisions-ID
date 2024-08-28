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
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.ListCityByIDRequest[[]int]]) (*[]entity.City, int64, error) {
			where := map[string]interface{}{}
			if ctx.Request.ID != nil {
				where["id"] = ctx.Request.ID
			}
			if ctx.Request.IDProvince != nil {
				where["id_province"] = ctx.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(ctx.DB, where)
			return &collections, total, err
		},
	)
}

func (uc *CityUseCase) GetById(ctx context.Context, request model.GetCityByIDRequest[[]int]) (*[]entity.City, int64, error) {
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetCityByIDRequest[[]int]]) (*[]entity.City, int64, error) {
			where := map[string]interface{}{}
			if ctx.Request.ID != nil {
				where["id"] = ctx.Request.ID
			}
			if ctx.Request.IDProvince != nil {
				where["id_province"] = ctx.Request.IDProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(ctx.DB, where)
			return &collections, total, err
		},
	)
}

func (uc *CityUseCase) GetFirstById(ctx context.Context, request model.GetCityByIDRequest[int]) (*entity.City, int64, error) {
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetCityByIDRequest[int]]) (*entity.City, int64, error) {
			id := ctx.Request.ID
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdProvince(ctx.DB, id, idProvince)

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
