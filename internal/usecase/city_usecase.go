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

type City struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.City[int, []int]
}

func NewCity(log *zap.Logger, db *gorm.DB, repository *repository.City[int, []int]) *City {
	return &City{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *City) List(ctx context.Context, request model.ListCityByIDRequest[[]int]) (*[]entity.City, int64, error) {
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.ListCityByIDRequest[[]int]]) (*[]entity.City, int64, error) {
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

func (uc *City) GetById(ctx context.Context, request model.GetCityByIDRequest[[]int]) (*[]entity.City, int64, error) {
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetCityByIDRequest[[]int]]) (*[]entity.City, int64, error) {
			id := ctx.Request.ID
			idProvince := ctx.Request.IDProvince

			where := map[string]interface{}{}
			if id != nil {
				where["id"] = id
			}
			if idProvince != nil {
				where["id_province"] = idProvince
			}

			collections, total, err := uc.Repository.FindAndCountBy(ctx.DB, where)

			if len(collections) == 0 {
				errorMessage := fmt.Sprintf("failed to get cities data with ID: %d", id)
				if idProvince != nil {
					errorMessage += fmt.Sprintf(" and ID Province: %d", idProvince)
				}
				return nil, 0, apperror.RecordNotFound(errorMessage)
			}

			return &collections, total, err
		},
	)
}

func (uc *City) GetFirstById(ctx context.Context, request model.GetCityByIDRequest[int]) (*entity.City, int64, error) {
	return appusecase.Wrapper[entity.City](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetCityByIDRequest[int]]) (*entity.City, int64, error) {
			id := ctx.Request.ID
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdProvince(ctx.DB, id, idProvince)

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0, apperror.RecordNotFound(fmt.Sprintf("failed to get cities data with ID: %d and ID Province: %d", id, idProvince))
			}

			return &collection, 1, err
		},
	)
}
