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

type Village struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.Village[int, []int]
}

func NewVillage(log *zap.Logger, db *gorm.DB, repository *repository.Village[int, []int]) *Village {
	return &Village{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *Village) List(ctx context.Context, request model.ListVillageByIDRequest[[]int]) (*[]entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.ListVillageByIDRequest[[]int]]) (*[]entity.Village, int64, error) {
			where := map[string]interface{}{}
			if ctx.Request.ID != nil {
				where["id"] = ctx.Request.ID
			}
			if ctx.Request.IDDistrict != nil {
				where["id_district"] = ctx.Request.IDDistrict
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

func (uc *Village) GetById(ctx context.Context, request model.GetVillageByIDRequest[[]int]) (*[]entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetVillageByIDRequest[[]int]]) (*[]entity.Village, int64, error) {
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

func (uc *Village) GetFirstById(ctx context.Context, request model.GetVillageByIDRequest[int]) (*entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.Context[model.GetVillageByIDRequest[int]]) (*entity.Village, int64, error) {
			id := ctx.Request.ID
			idDistrict := ctx.Request.IDDistrict
			idCity := ctx.Request.IDCity
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdDistrictAndIdCityAndIdProvince(ctx.DB, id, idDistrict, idCity, idProvince)

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0, apperror.RecordNotFound(fmt.Sprintf("failed to get cities data with ID: %d and ID District: %d and ID City: %d and ID Province: %d", id, idDistrict, idCity, idProvince))
			}

			return &collection, 1, err
		},
	)
}
