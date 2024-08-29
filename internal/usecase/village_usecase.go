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

type VillageUseCase struct {
	Log        *zap.Logger
	DB         *gorm.DB
	Repository repository.VillageRepository[int, []int]
}

func NewVillageUseCase(log *zap.Logger, db *gorm.DB, repository *repository.VillageRepository[int, []int]) *VillageUseCase {
	return &VillageUseCase{
		Log:        log,
		DB:         db,
		Repository: *repository,
	}
}

func (uc *VillageUseCase) List(ctx context.Context, request model.ListVillageByIDRequest[[]int]) (*[]entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.ListVillageByIDRequest[[]int]]) (*[]entity.Village, int64, error) {
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

func (uc *VillageUseCase) GetById(ctx context.Context, request model.GetVillageByIDRequest[[]int]) (*[]entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetVillageByIDRequest[[]int]]) (*[]entity.Village, int64, error) {
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

func (uc *VillageUseCase) GetFirstById(ctx context.Context, request model.GetVillageByIDRequest[int]) (*entity.Village, int64, error) {
	return appusecase.Wrapper[entity.Village](
		appusecase.NewContext(ctx, uc.Log, uc.DB, request),
		func(ctx *appusecase.UseCaseContext[model.GetVillageByIDRequest[int]]) (*entity.Village, int64, error) {
			id := ctx.Request.ID
			idDistrict := ctx.Request.IDDistrict
			idCity := ctx.Request.IDCity
			idProvince := ctx.Request.IDProvince
			collection, err := uc.Repository.FirstByIdAndIdDistrictAndIdCityAndIdProvince(ctx.DB, id, idDistrict, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d and ID District: %d and ID City: %d and ID Province: %d", id, idDistrict, idCity, idProvince)
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
