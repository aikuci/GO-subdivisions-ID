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
	CrudUseCase CrudUseCase[entity.District] // embedded

	Repository repository.DistrictRepository[int, []int]
}

func NewDistrictUseCase(logger *zap.Logger, db *gorm.DB, repository *repository.DistrictRepository[int, []int]) *DistrictUseCase {
	crudUseCase := NewCrudUseCase(logger, db, repository)

	return &DistrictUseCase{
		CrudUseCase: *crudUseCase,

		Repository: *repository,
	}
}

func (uc *DistrictUseCase) List(ctx context.Context, request model.ListRequest) ([]entity.District, error) {
	return uc.CrudUseCase.List(ctx, request)
}
func (uc *DistrictUseCase) GetById(ctx context.Context, request model.GetByIDRequest[int]) ([]entity.District, error) {
	return uc.CrudUseCase.GetById(ctx, request)
}
func (uc *DistrictUseCase) GetByIds(ctx context.Context, request model.GetByIDRequest[[]int]) ([]entity.District, error) {
	return uc.CrudUseCase.GetByIds(ctx, request)
}
func (uc *DistrictUseCase) GetFirstById(ctx context.Context, request model.GetByIDRequest[int]) (*entity.District, error) {
	return uc.CrudUseCase.GetFirstById(ctx, request)
}

// Specific UseCase
func (uc *DistrictUseCase) ListFindByIdAndIdCityAndIdProvince(ctx context.Context, request model.ListDistrictByIDRequest[[]int]) ([]entity.District, error) {
	useCase := newUseCase[entity.District](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperPlural(
		ctx,
		useCase,
		func(cp *CallbackParam[model.ListDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			where := map[string]interface{}{}
			if cp.request.ID != nil {
				where["id"] = cp.request.ID
			}
			if cp.request.IDCity != nil {
				where["id_city"] = cp.request.IDCity
			}
			if cp.request.IDProvince != nil {
				where["id_province"] = cp.request.IDProvince
			}
			return uc.Repository.FindBy(cp.tx, where)
		},
	)
}

func (uc *DistrictUseCase) GetFindByIdAndIdCityAndIdProvince(ctx context.Context, request model.GetDistrictByIDRequest[[]int]) ([]entity.District, error) {
	useCase := newUseCase[entity.District](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperPlural(
		ctx,
		useCase,
		func(cp *CallbackParam[model.GetDistrictByIDRequest[[]int]]) ([]entity.District, error) {
			where := map[string]interface{}{}
			if cp.request.ID != nil {
				where["id"] = cp.request.ID
			}
			if cp.request.IDCity != nil {
				where["id_city"] = cp.request.IDCity
			}
			if cp.request.IDProvince != nil {
				where["id_province"] = cp.request.IDProvince
			}
			return uc.Repository.FindBy(cp.tx, where)
		},
	)
}

func (uc *DistrictUseCase) GetFirstByIdAndIdCityAndIdProvince(ctx context.Context, request model.GetDistrictByIDRequest[int]) (*entity.District, error) {
	useCase := newUseCase[entity.District](uc.CrudUseCase.Log, uc.CrudUseCase.DB, request)

	return wrapperSingular(
		ctx,
		useCase,
		func(cp *CallbackParam[model.GetDistrictByIDRequest[int]]) (*entity.District, error) {
			id := cp.request.ID
			idProvince := cp.request.IDProvince
			idCity := cp.request.IDCity
			collection, err := uc.Repository.FirstByIdAndIdCityAndIdProvince(cp.tx, id, idCity, idProvince)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					errorMessage := fmt.Sprintf("failed to get cities data with ID: %d, ID City: %v and ID Province: %d", id, idCity, idProvince)
					cp.log.Warn(err.Error(), zap.String("errorMessage", errorMessage))
					return nil, apperror.RecordNotFound(errorMessage)
				}

				cp.log.Warn(err.Error())
				return nil, err
			}

			return collection, nil
		},
	)
}
