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

func (uc *DistrictUseCase) List(ctx context.Context, request model.ListDistrictByIDRequest[[]int]) ([]entity.District, error) {
	return wrapperPlural(
		newUseCase[entity.District](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.ListDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			where := map[string]interface{}{}
			if ca.request.ID != nil {
				where["id"] = ca.request.ID
			}
			if ca.request.IDCity != nil {
				where["id_city"] = ca.request.IDCity
			}
			if ca.request.IDProvince != nil {
				where["id_province"] = ca.request.IDProvince
			}
			return uc.Repository.FindBy(ca.tx, where)
		},
	)
}
func (uc *DistrictUseCase) GetById(ctx context.Context, request model.GetDistrictByIDRequest[[]int]) ([]entity.District, error) {
	return wrapperPlural(
		newUseCase[entity.District](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			where := map[string]interface{}{}
			if ca.request.ID != nil {
				where["id"] = ca.request.ID
			}
			if ca.request.IDCity != nil {
				where["id_city"] = ca.request.IDCity
			}
			if ca.request.IDProvince != nil {
				where["id_province"] = ca.request.IDProvince
			}
			return uc.Repository.FindBy(ca.tx, where)
		},
	)
}
func (uc *DistrictUseCase) GetFirstById(ctx context.Context, request model.GetDistrictByIDRequest[int]) (*entity.District, error) {
	return wrapperSingular(
		newUseCase[entity.District](ctx, uc.Log, uc.DB, request),
		func(ca *CallbackArgs[model.GetDistrictByIDRequest[int]]) (*entity.District, error) {
			id := ca.request.ID
			idProvince := ca.request.IDProvince
			idCity := ca.request.IDCity
			collection, err := uc.Repository.FirstByIdAndIdCityAndIdProvince(ca.tx, id, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d, ID City: %v and ID Province: %d", id, idCity, idProvince)
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
