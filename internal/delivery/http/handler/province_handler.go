package handler

import (
	"github.com/aikuci/go-subdivisions-id/internal/entity"
	"github.com/aikuci/go-subdivisions-id/internal/model"
	"github.com/aikuci/go-subdivisions-id/internal/model/mapper"
	apphandler "github.com/aikuci/go-subdivisions-id/pkg/delivery/http/handler"
	appusecase "github.com/aikuci/go-subdivisions-id/pkg/usecase"

	"go.uber.org/zap"
)

type Province struct {
	CrudHandler *apphandler.Crud[entity.Province, model.ProvinceResponse]
}

func NewProvince(log *zap.Logger, useCase appusecase.CruderUseCase[entity.Province]) *Province {
	return &Province{
		CrudHandler: apphandler.NewCrud(log, useCase, mapper.NewProvince()),
	}
}
